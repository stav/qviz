package routes

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

type Quiz struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"created_at"`
	Ques []Question `json:"question"`
	Msg  string
}

type Question struct {
	// ID   int    `json:"id"`
	Text string `json:"text"`
}

func AppIndexHandler(c echo.Context) error {
	var results []Quiz
	err := supabase.DB.From("quiz").Select("*").Execute(&results)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.Render(200, "app", results)
}

func AppQuizHandler(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("ID:", id)

	var results Quiz
	var quiz Quiz
	var err error

	// First query to get the quiz
	err = supabase.DB.From("quiz").Select("*").Single().Filter("id", "eq", id).Execute(&quiz)
	if err != nil {
		return c.JSON(500, err)
	}
	fmt.Println("Quiz:", quiz)

	// Second query to get the questions
	err = supabase.DB.From("quiz").Select("*,question!inner(text)").Single().Filter("id", "eq", id).Execute(&results)
	if err != nil {
		fmt.Println("Error:", err)
		// Check if there are no questions
		if strings.HasPrefix(err.Error(), "PGRST116") {
			quiz.Msg = err.Error()
			return c.Render(200, "quiz.html", quiz)
		}
		return c.JSON(500, err)
	}
	fmt.Println("Quizs:", results)

	quiz.Ques = make([]Question, len(results.Ques))
	for i, q := range results.Ques {
		fmt.Println("iQ:", i, q)
		quiz.Ques[i] = q
	}
	fmt.Println("Quiz2:", quiz)

	return c.Render(200, "quiz.html", quiz)
}

// supabase.DB.From(table).Select("item,customer!inner(id, short_name, full_name)").Filter(`customer.id`, "eq", "0")
