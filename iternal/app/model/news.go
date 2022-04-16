package model

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type News struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo
	TimeDate string `json:"time-date"`
}

type Photo struct {
	Payload     io.Reader
	PayloadName string `json:"payload-name"`
	PayloadSize int64  `json:"payload-size"`
}

func GenerateObjectName(news *News) string {
	t := time.Now()
	formatted := fmt.Sprintf(
		"%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(),
	)
	return fmt.Sprintf("%d/%s.%s", rand.Intn(100), formatted, "png")
}
