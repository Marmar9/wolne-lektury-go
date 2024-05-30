package controllers

import (
	"github.com/Marmar9/wolne-lektury-go/server/db/models"
	"github.com/Marmar9/wolne-lektury-go/server/repositories"
)

type GetLecturesController struct {
	repository *repositories.Repository
}

func (this *GetLecturesController) GetLectures(authorName string, kind string) []models.Lecture {
	return this.repository.GetLectures(authorName, kind)
}

func NewGetLecturesController(repo *repositories.Repository) *GetLecturesController {
	return &GetLecturesController{repo}
}
