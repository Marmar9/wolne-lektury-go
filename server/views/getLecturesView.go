package views

import (
	"net/http"

	"github.com/Marmar9/wolne-lektury-go/server/controllers"
	"github.com/Marmar9/wolne-lektury-go/server/utils"
	"github.com/labstack/echo/v4"
)

func GetGetLecturesView(controller *controllers.GetLecturesController) echo.HandlerFunc {

	return func(c echo.Context) error {
		var body utils.GetLecturesBody

		err := c.Bind(&body)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		response := controller.GetLectures(body.AuthorName, body.Kind)

		return c.JSON(http.StatusOK, response)
	}
}
