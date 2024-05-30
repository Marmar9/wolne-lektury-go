package utils

import (
	"os"

	"github.com/Marmar9/wolne-lektury-go/server/db/models"
	"github.com/joho/godotenv"
)

func GetEnv() (string, string, string) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	return dbUser, dbPass, dbName
}

type AddLectureBody struct {
	Lecture    models.Lecture `json:"lecture"`
	Kind       string         `json:"kind"`
	AuthorName string         `json:"authorName"`
}

type GetLecturesBody struct {
	AuthorName string `query:"authorName"`
	Kind       string `query:"kind"`
}
