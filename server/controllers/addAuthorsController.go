package controllers

import (
	"github.com/Marmar9/wolne-lektury-go/server/db/models"
	"github.com/Marmar9/wolne-lektury-go/server/repositories"
)

type AddAuthorsController struct {
	repository *repositories.Repository
}

func (this *AddAuthorsController) AddAuthors(authors []models.Author) {
	this.repository.AddAuthors(authors)
}

func NewAddAuthorsController(repo *repositories.Repository) *AddAuthorsController {
	return &AddAuthorsController{repo}
}
