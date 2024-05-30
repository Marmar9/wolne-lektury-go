package views

import (
	"fmt"
	"net/http"

	"github.com/Marmar9/wolne-lektury-go/server/controllers"
	"github.com/Marmar9/wolne-lektury-go/server/utils"
	"github.com/labstack/echo/v4"
)

func GetAddLectureView(controller *controllers.AddLectureController) echo.HandlerFunc {
	return func(c echo.Context) error {
		var body utils.AddLectureBody

		err := c.Bind(&body)
		fmt.Println(body)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		controller.AddLecture(body)
		return c.String(http.StatusOK, "created lecture")
	}
}
