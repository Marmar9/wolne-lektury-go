package controllers

import (
	"github.com/Marmar9/wolne-lektury-go/server/db/models"
	"github.com/Marmar9/wolne-lektury-go/server/repositories"
)

type AddKindsController struct {
	repository *repositories.Repository
}

func (this *AddKindsController) AddKinds(kinds []models.Kind) {
	this.repository.AddKinds(kinds)
}

func NewAddKindsController(repo *repositories.Repository) *AddKindsController {
	return &AddKindsController{repo}
}
