package router

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"

	"github.com/labstack/echo/v4"
)

func QuizFromId(quiz_id, question_num string) Quiz {
	fmt.Println("IDs:", quiz_id, question_num)

	var err error
	var quiz Quiz
	var questions []Question

	// First query to get the quiz
	// We should only need one db call but I can't figure out how to outer join
	// the questions and answers in one query with nedpals/supabase-go library
	err = supabase.DB.From("quiz").Select("*").Single().Filter("id", "eq", quiz_id).Execute(&quiz)
	if err != nil {
		quiz.Msg = err.Error()
		return quiz
	}
	fmt.Println("Quiz:", quiz)

	// Second query to get the questions & answers
	query := "id,number,text,answer!left(id,text,is_correct)"
	err = supabase.DB.From("question").Select(query).Filter("quiz_id", "eq", quiz_id).Execute(&questions)
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
	quiz.Qcnt = len(questions)
	if quiz.Qcnt == 0 {
		quiz.Msg = "No questions found"
		return quiz
	}

	// Now filter the questions to get the one with the question number.
	// I can filter the questions with the client but I don't know how to also
	// `SELECT COUNT(*)` in the same query. I could make a separate db call to
	// get the count but I decided to just return the data and filter it here.
	filtered_questions := funk.Filter(questions, func(q Question) bool {
		num, _ := strconv.Atoi(question_num)
		return q.Num == num
	}).([]Question)
	fmt.Println("filtered_questions:", filtered_questions)
	quiz.Ques = filtered_questions[0]
	fmt.Println("Quiz2:", quiz)

	// Return the quiz
	return quiz
}

func ApiQuizHandler(c echo.Context) error {
	quiz_id := c.Param("quizId")
	fmt.Println("ApiQuizHandler Quiz ID:", quiz_id)
	return c.Render(200, "__quiz.html", QuizFromId(quiz_id, "1"))
}

func ApiQuestionHandler(c echo.Context) error {
	quiz_id := c.Param("quizId")
	question_num := c.Param("questionNumber")
	fmt.Println("ApiQuestionHandler:", quiz_id, question_num)
	return c.Render(200, "__quiz.html", QuizFromId(quiz_id, question_num))
}

func ApiAnswerHandler(c echo.Context) error {
	quiz_id := c.Param("quizId")
	question_num := c.Param("questionNumber")
	answer_id := c.Param("answerId")
	fmt.Println("ApiAnswerHandler:", quiz_id, question_num, answer_id)
	return c.Render(200, "__answer.html", answer_id)
}
