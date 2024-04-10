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
	ID   int    `json:"id"`
	Text string `json:"text"`
	Ans  []Answer `json:"answer"`
}

type Answer struct {
	ID	 int    `json:"id"`
	Text string `json:"text"`
	True bool   `json:"is_correct"`
}

func AppIndexHandler(c echo.Context) error {
	var quizs []Quiz
	err := supabase.DB.From("quiz").Select("*").Execute(&quizs)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.Render(200, "app", quizs)
}

func AppQuizHandler(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("ID:", id)

	var quizQs Quiz
	var quiz Quiz
	var err error

	// First query to get the quiz
	err = supabase.DB.From("quiz").Select("*").Single().Filter("id", "eq", id).Execute(&quiz)
	if err != nil {
		return c.JSON(500, err)
	}
	fmt.Println("Quiz:", quiz)

	// Second query to get the questions & answers
	err = supabase.DB.From("quiz").Select("*,question!inner(id,text,answer!left(id,text,is_correct))").Single().Filter("id", "eq", id).Execute(&quizQs)
	if err != nil {
		fmt.Println("Error:", err)
		// Check if there are no questions
		if strings.HasPrefix(err.Error(), "PGRST116") {
			quiz.Msg = err.Error()
			return c.Render(200, "__quiz.html", quiz)
		}
		return c.JSON(500, err)
	}
	fmt.Println("quizQs:", quizQs)

	// Copy the questions to the quiz one at a time because of the different types
	quiz.Ques = make([]Question, len(quizQs.Ques))
	for i, q := range quizQs.Ques {
		fmt.Println("iQ:", i, q)
		quiz.Ques[i] = q
	}
	fmt.Println("Quiz2:", quiz)

	return c.Render(200, "__quiz.html", quiz)
}
