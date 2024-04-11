package routes

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

func QuizFromId(c echo.Context) Quiz {
	id := c.Param("id")
	fmt.Println("ID:", id)

	var quizQs Quiz
	var quiz Quiz
	var err error

	// First query to get the quiz
	// We only need this first call because I can't figure out how to outer join
	// the questions and answers in the second call which should be the only call
	err = supabase.DB.From("quiz").Select("*").Single().Filter("id", "eq", id).Execute(&quiz)
	if err != nil {
		quiz.Msg = err.Error()
		return quiz
	}
	fmt.Println("Quiz:", quiz)

	// Second query to get the questions & answers
	err = supabase.DB.From("quiz").Select("*,question!inner(id,text,answer!left(id,text,is_correct))").Single().Filter("id", "eq", id).Execute(&quizQs)
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
	fmt.Println("quizQs:", quizQs)

	// Copy the questions to the quiz one at a time because of the different types
	quiz.Ques = make([]Question, len(quizQs.Ques))
	for i, q := range quizQs.Ques {
		fmt.Println("iQ:", i, q)
		quiz.Ques[i] = q
	}
	fmt.Println("Quiz2:", quiz)

	return quiz
}

func ApiQuizHandler(c echo.Context) error {
	return c.Render(200, "__quiz.html", QuizFromId(c))
}
