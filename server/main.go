package main

import (
	"github.com/Marmar9/wolne-lektury-go/server/controllers"
	"github.com/Marmar9/wolne-lektury-go/server/repositories"
	"github.com/Marmar9/wolne-lektury-go/server/views"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	repository := repositories.NewRepository()

	e.POST("/lecture", views.GetAddLectureView(controllers.NewAddLectureController(repository)))
	e.GET("/lectures", views.GetGetLecturesView(controllers.NewGetLecturesController(repository)))
	e.POST("/categories/authors", views.GetAddAuthorsView(controllers.NewAddAuthorsController(repository)))
	e.POST("/categories/kinds", views.GetAddKindsView(controllers.NewAddKindsController(repository)))

	e.Logger.Fatal(e.Start(":8080"))
}
