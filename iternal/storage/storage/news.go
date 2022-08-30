package storage

import (
	"CIS_Backend_Server/iternal/model"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"net/url"
	"time"
)

type NewsRepository struct {
	storage *Storage
}

func (r *NewsRepository) Create(ctx context.Context, n *model.News) error {
	//creating a unique photo name
	nameGenerator(n)

	//sending photos to minio
	_, err := r.storage.minioClient.PutObject(
		ctx,
		r.storage.bucketName,
		n.NameWithUUID,
		n.Payload,
		n.Size,
		minio.PutObjectOptions{ContentType: n.ContentType},
	)
	if err != nil {
		return err
	}

	//sending to the database news content
	row := r.storage.db.QueryRow(
		"INSERT INTO news (title, description, photo) VALUES ($1, $2, $3) RETURNING id, time_date",
		n.Title,
		n.Description,
		n.NameWithUUID,
	)

	if err := row.Err(); err != nil {
		//an error occurred in the database, deleting photos from minio
		secondErr := r.storage.minioClient.RemoveObject(ctx, r.storage.bucketName, n.NameWithUUID, minio.RemoveObjectOptions{})

		//error when rolling back changes from minio, returning both errors from database and minio
		if secondErr != nil {
			return errors.New("first error: " + err.Error() + " | second error: " + secondErr.Error())
		}

		return err
	}

	//getting id and date with time from database
	err = row.Scan(&n.Id, &n.TimeDate)

	return err
}

func (r *NewsRepository) Get(ctx context.Context, id int) (news []model.News, err error) {
	var rows *sql.Rows

	if id == 1 {
		//getting the latest 12 news from the database
		rows, err = r.storage.db.Query("SELECT * FROM news ORDER BY id DESC LIMIT 12")
	} else {
		//receiving follow-up news by latest news id from the database
		rows, err = r.storage.db.Query("SELECT * FROM news WHERE id < $1 ORDER BY id DESC LIMIT 12", id)
	}

	//checking for data in the database
	if errors.Is(err, sql.ErrNoRows) {
		return nil, model.ErrNewsNotFound
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := model.News{}

		//getting id, title, description, photo title and date with time from database
		err := rows.Scan(&n.Id, &n.Title, &n.Description, &n.NameWithUUID, &n.TimeDate)
		if err != nil {
			logrus.Error(err)
			continue
		}

		//setting request parameters for content
		reqParams := make(url.Values)
		reqParams.Set(
			"response-content-disposition",
			"attachment; filename=\""+n.NameWithUUID+"\"",
		)

		//generates a presigned url which expires in a day
		photoURL, err := r.storage.minioClient.PresignedGetObject(
			ctx,
			r.storage.bucketName,
			n.NameWithUUID,
			time.Hour*24,
			reqParams,
		)
		if err != nil {
			return nil, err
		}

		//url type to string type
		n.URL = photoURL.String()

		news = append(news, n)
	}

	return news, err
}

func (r *NewsRepository) Change(ctx context.Context, n *model.News) error {
	//checking the existence of the updated news in the database
	err := r.storage.db.QueryRow("SELECT photo FROM news WHERE id = $1", n.Id).Scan(&n.NameWithUUID)
	if errors.Is(err, sql.ErrNoRows) {
		return model.ErrNewsNotFound
	}
	if err != nil {
		return err
	}

	//creating a unique photo name
	nameGenerator(n)

	//changing the news in the database according to new data
	_, err = r.storage.db.Exec("UPDATE news SET title = $1, description = $2, photo = $3 WHERE id = $4",
		n.Title,
		n.Description,
		n.NameWithUUID,
		n.Id,
	)
	if err != nil {
		return err
	}

	//adding a new photo news
	_, err = r.storage.minioClient.PutObject(
		ctx,
		r.storage.bucketName,
		n.NameWithUUID,
		n.Payload,
		n.Size,
		minio.PutObjectOptions{ContentType: n.ContentType},
	)

	if err != nil {
		//rolling back changes to the database
		_, secondErr := r.storage.db.Exec("DELETE FROM news WHERE id = $1", n.Id)

		//checking for database errors when reverting changes, returning two errors from minio and from the database
		if secondErr != nil {
			return errors.New("first error: " + err.Error() + " | second error: " + secondErr.Error())
		}

		return err
	}

	//delete old news photo
	err = r.storage.minioClient.RemoveObject(ctx, r.storage.bucketName, n.NameWithUUID, minio.RemoveObjectOptions{})

	return err
}

func (r *NewsRepository) Delete(ctx context.Context, id int) error {
	var name string

	//getting photo name by id news
	err := r.storage.db.QueryRow("SELECT photo FROM news WHERE id = $1", id).Scan(&name)
	if errors.Is(err, sql.ErrNoRows) {
		return model.ErrNewsNotFound
	}
	if err != nil {
		return err
	}

	//deleting news from the database
	_, err = r.storage.db.Exec("DELETE FROM news WHERE id = $1", id)
	if err != nil {
		return err
	}

	//deleting news from minio
	err = r.storage.minioClient.RemoveObject(ctx, r.storage.bucketName, name, minio.RemoveObjectOptions{})

	return err
}

//nameGenerator generating a unique photo name by embedding an uuid before the file extension
func nameGenerator(n *model.News) {
	if string(n.Name[len(n.Name)-5]) == "." {
		n.NameWithUUID = n.Name[:len(n.Name)-4] + uuid.NewString() + n.Name[len(n.Name)-5:]
	} else {
		n.NameWithUUID = n.Name[:len(n.Name)-3] + uuid.NewString() + n.Name[len(n.Name)-4:]
	}
}
