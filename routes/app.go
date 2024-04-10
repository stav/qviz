package routes

import (
	"github.com/labstack/echo/v4"
)

func AppIndexHandler(c echo.Context) error {
	var quizs []Quiz
	err := supabase.DB.From("quiz").Select("*").Execute(&quizs)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.Render(200, "app.html", quizs)
}

func AppQuizHandler(c echo.Context) error {
	return c.Render(200, "quiz.html", QuizFromId(c))
}
