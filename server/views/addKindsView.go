package views

import (
	"fmt"
	"net/http"

	"github.com/Marmar9/wolne-lektury-go/server/controllers"
	"github.com/Marmar9/wolne-lektury-go/server/db/models"
	"github.com/labstack/echo/v4"
)

func GetAddKindsView(controller *controllers.AddKindsController) echo.HandlerFunc {

	return func(c echo.Context) error {
		var body []models.Kind

		err := c.Bind(&body)
		fmt.Println(body)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		controller.AddKinds(body)
		return c.String(http.StatusOK, "created kinds")
	}
}
