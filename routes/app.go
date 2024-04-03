package routes

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AppIndexHandler(c echo.Context) error {
  fmt.Println("GET /app")
  return c.JSON(http.StatusOK, "{ \"message\": \"GET /app\" }")
}
