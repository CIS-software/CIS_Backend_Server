package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"context"
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

func (r *NewsRepository) CreateNews(ctx context.Context, n *model.News) error {
	n.Name = fmt.Sprintf("%s.%s", uuid.NewString(), "png")
	_, err := r.storage.minioClient.PutObject(
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

	return r.storage.db.QueryRow(
		"INSERT INTO news (title, description, photo) VALUES ($1, $2, $3) RETURNING id, time_date",
		n.Title,
		n.Description,
		n.Name,
	).Scan(&n.Id, &n.TimeDate)
}

func (r *NewsRepository) GetNews(ctx context.Context) (news []model.News, err error) {
	rows, err := r.storage.db.Query("SELECT * FROM news ORDER BY time_date DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		n := model.News{}
		err := rows.Scan(&n.Id, &n.Title, &n.Description, &n.Name, &n.TimeDate)
		if err != nil {
			logrus.Error(err)
			continue
		}
		reqParams := make(url.Values)
		reqParams.Set(
			"response-content-disposition",
			"attachment; filename=\""+n.Name+"\"",
		)
		photoURL, err := r.storage.minioClient.PresignedGetObject(
			ctx,
			r.storage.bucketName,
			n.Name,
			time.Hour,
			reqParams,
		)
		if err != nil {
			return nil, err
		}
		n.URL = photoURL.String()
		news = append(news, n)
	}

	return news, err
}

func (r *NewsRepository) UpdateNews(ctx context.Context, n *model.News) error {
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

func (r *NewsRepository) DeleteNews(ctx context.Context, id int) error {
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
