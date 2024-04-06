package routes

import (
	"github.com/labstack/echo/v4"
)

type Quiz struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"created_at"`
}

func AppIndexHandler(c echo.Context) error {
	var results []Quiz
	err := supabase.DB.From("quiz").Select("*").Execute(&results)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.Render(200, "app", results)
}
