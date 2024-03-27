package routes

import (
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.Render(200, "index", nil)
}

func QuizHandler(c echo.Context) error {
	return c.Render(200, "quiz", nil)
}
