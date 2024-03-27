package routes

import (
	"github.com/labstack/echo/v4"
)

func IndexHandler(c echo.Context) error {
  return c.String(200, "Hello, World!!\n")
}
