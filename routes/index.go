package routes

import (
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
	return c.Render(200, "index", nil)
}

func LoginHandler(c echo.Context) error {
	return c.String(200, "Auth attempt")
}
