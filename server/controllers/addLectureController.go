package controllers

import (
	"github.com/Marmar9/wolne-lektury-go/server/repositories"
	"github.com/Marmar9/wolne-lektury-go/server/utils"
)

type AddLectureController struct {
	repository *repositories.Repository
}

func (this *AddLectureController) AddLecture(lectureBody utils.AddLectureBody) {
	this.repository.AddLecture(lectureBody)
}

func NewAddLectureController(repo *repositories.Repository) *AddLectureController {
	return &AddLectureController{repo}
}
