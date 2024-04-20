package router

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func QuizFromId(quiz_id, quest_num string) Quiz {

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

	// Second query to get the questions & answers
	query := "id,number,text,answer!left(id,text,is_correct)"
	err = supabase.DB.From("questions").Select(query).Filter("quiz_id", "eq", quiz_id).Execute(&questions)
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

	// Copy the questions to the quiz one at a time because of the different types
	quiz.Questions = make([]Question, len(questions))
	for i, q := range questions {
		fmt.Println("iQ:", i, q)
		quiz.Questions[i] = q
	}
	quiz.Qcount = len(quiz.Questions)
	if quiz.Qcount == 0 {
		quiz.Msg = "No questions found"
		return quiz
	}

	// Set the question to the one requested
	index, err := strconv.Atoi(quest_num) // Convert quest_num from string to int
	if err != nil {
		quiz.Msg = err.Error()
		fmt.Println("Error:", err)
		return quiz
	}
	if index < 1 || index > len(quiz.Questions) {
		quiz.Msg = "Question number out of range"
		fmt.Println("Error:", quiz.Msg)
		return quiz
	}
	index-- // decrement index to get the correct index
	quiz.Questions = []Question{quiz.Questions[index]}
	fmt.Println("Quiz:", index, quiz)

	return quiz
}

func ApiQuizHandler(c echo.Context) error {
	quiz_id := c.Param("quizId")
	return c.Render(200, "__quiz.html", QuizFromId(quiz_id, "1"))
}

func ApiQuestionHandler(c echo.Context) error {
	quiz_id := c.Param("quizId")
	quest_num := c.Param("questionNumber")
	return c.Render(200, "__quiz.html", QuizFromId(quiz_id, quest_num))
}

func ApiAnswerHandler(c echo.Context) error {
	quiz_id := c.Param("quizId")
	quest_num := c.Param("questionNumber")
	answer_id := c.Param("answerId")
	fmt.Println("ApiAnswerHandler:", quiz_id, quest_num, answer_id)
	return c.Render(200, "__answer.html", answer_id)
}
