package dbstorage

import (
	"CIS_Backend_Server/iternal/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	//adds the last element of the slice from the previous last element
	//then the penultimate element is replaced by a new uuid
	n.NameSlice = append(n.NameSlice, n.NameSlice[len(n.NameSlice)-1])
	n.NameSlice[len(n.NameSlice)-2] = uuid.NewString()

	//getting a photo name string from a slice, keeping all the dots
	for index := range n.NameSlice {
		//check for the last element, to skip the dot
		if len(n.NameSlice) == index+1 {
			n.Name = n.Name + n.NameSlice[index]
			break
		}
		n.Name = n.Name + n.NameSlice[index] + "."
	}

	//sending photos to minio
	_, err := r.storage.minioClient.PutObject(
		ctx,
		r.storage.bucketName,
		n.Name,
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
		n.Name,
	)
	if err := row.Err(); err != nil {
		return err
	}

	//getting id and date with time from database
	err = row.Scan(&n.Id, &n.TimeDate)

	return err
}

func (r *NewsRepository) Get(ctx context.Context) (news []model.News, err error) {
	//getting all the news from the database
	rows, err := r.storage.db.Query("SELECT * FROM news ORDER BY time_date DESC")

	//checking for data in the database
	if errors.Is(err, sql.ErrNoRows) {
		return nil, sql.ErrNoRows
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := model.News{}

		//getting id, title, description, photo title and date with time from database
		err := rows.Scan(&n.Id, &n.Title, &n.Description, &n.Name, &n.TimeDate)
		if err != nil {
			logrus.Error(err)
			continue
		}

		//setting request parameters for content
		reqParams := make(url.Values)
		reqParams.Set(
			"response-content-disposition",
			"attachment; filename=\""+n.Name+"\"",
		)

		//generates a presigned url which expires in a day
		photoURL, err := r.storage.minioClient.PresignedGetObject(
			ctx,
			r.storage.bucketName,
			n.Name,
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
	err := r.storage.db.QueryRow("SELECT photo FROM news WHERE id = $1", n.Id).Scan(&n.Name)
	if err != nil {
		return err
	}
	name := n.Name
	n.Name = fmt.Sprintf("%s.%s", uuid.NewString(), "png")

	result, err := r.storage.db.Exec("UPDATE news SET title = $1, description = $2, photo = $3 WHERE id = $4",
		n.Title,
		n.Description,
		n.Name,
		n.Id,
	)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("news not found")
	}

	_, err = r.storage.minioClient.PutObject(
		ctx,
		r.storage.bucketName,
		n.Name,
		n.Payload,
		n.Size,
		minio.PutObjectOptions{ContentType: "image/png"},
	)
	if err != nil {
		return err
	}

	err = r.storage.minioClient.RemoveObject(ctx, r.storage.bucketName, name, minio.RemoveObjectOptions{})
	return err
}

func (r *NewsRepository) Delete(ctx context.Context, id int) error {
	var name string
	err := r.storage.db.QueryRow("SELECT photo FROM news WHERE id = $1", id).Scan(&name)
	if err != nil {
		return err
	}

	result, err := r.storage.db.Exec("DELETE FROM news WHERE id = $1", id)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("news not found")
	}

	err = r.storage.minioClient.RemoveObject(ctx, r.storage.bucketName, name, minio.RemoveObjectOptions{})
	return err
}
