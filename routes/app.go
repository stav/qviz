package routes

import (
	"github.com/labstack/echo/v4"
)

func AppIndexHandler(c echo.Context) error {
	return c.Render(200, "app", "Qviz App")
}
