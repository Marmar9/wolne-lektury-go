package repositories

import (
	"fmt"

	"github.com/Marmar9/wolne-lektury-go/server/db/models"
	"github.com/Marmar9/wolne-lektury-go/server/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func (this *Repository) GetLectures(authorName string, kind string) []models.Lecture {
	var lectures []models.Lecture
	query := this.db.Preload("Author").Preload("Kind")

	// Then, filter lectures by author name and kind name
	query = query.Select("lectures.title, lectures.content, lectures.author_id, lectures.kind_id").
		Joins("JOIN authors ON lectures.author_id = authors.id").
		Joins("JOIN kinds ON lectures.kind_id = kinds.id").
		Where("authors.name = ? AND kinds.name = ?", authorName, kind).
		Find(&lectures)
	return lectures
}

func (this *Repository) AddLecture(lectureBody utils.AddLectureBody) {

	var author models.Author
	this.db.Where("name = ?", lectureBody.AuthorName).First(&author)
	lectureBody.Lecture.AuthorID = author.ID

	var kind models.Kind
	this.db.Where("name = ?", lectureBody.Kind).First(&kind)
	lectureBody.Lecture.KindID = kind.ID

	this.db.Create(&lectureBody.Lecture)
}

func (this *Repository) AddAuthors(authors []models.Author) {
	this.db.Create(&authors)
}

func (this *Repository) AddKinds(kinds []models.Kind) {
	this.db.Create(&kinds)
}

func NewRepository() *Repository {
	dbUser, dbPass, dbName := utils.GetEnv()
	connectionString := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable", dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &Repository{db}
}
