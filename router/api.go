package router

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func QuizFromId(c echo.Context) Quiz {
	id := c.Param("id")
	fmt.Println("ID:", id)

	var err error
	var quiz Quiz
	var questions []Question

	// First query to get the quiz
	// We should only need one db call but I can't figure out how to outer join
	// the questions and answers in one query with nedpals/supabase-go library
	err = supabase.DB.From("quiz").Select("*").Single().Filter("id", "eq", id).Execute(&quiz)
	if err != nil {
		quiz.Msg = err.Error()
		return quiz
	}
	fmt.Println("Quiz:", quiz)

	// Second query to get the questions & answers
	query := "id,text,answer!left(id,text,is_correct)"
	err = supabase.DB.From("question").Select(query).Filter("quiz_id", "eq", id).Execute(&questions)
	if err != nil {
		fmt.Println("Error:", err)
		// Check if there are no questions
		if strings.HasPrefix(err.Error(), "PGRST116") {
			quiz.Msg = err.Error()
			return quiz
		}
		quiz.Msg = err.Error()
		return quiz
	}
	fmt.Println("questions:", questions)

	// Copy the questions to the quiz one at a time because of the different types
	quiz.Ques = make([]Question, len(questions))
	for i, q := range questions {
		fmt.Println("iQ:", i, q)
		quiz.Ques[i] = q
	}
	fmt.Println("Quiz2:", quiz)

	return quiz
}

func ApiQuizHandler(c echo.Context) error {
	return c.Render(200, "__quiz.html", QuizFromId(c))
}
