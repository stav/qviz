package routes

import (
	"os"

	"github.com/labstack/echo/v4"

	supa "github.com/nedpals/supabase-go"

	_ "github.com/joho/godotenv/autoload"
)

var SUPABASE_URL string = os.Getenv("SUPABASE_URL")
var SUPABASE_KEY string = os.Getenv("SUPABASE_KEY")
var supabase = supa.CreateClient(SUPABASE_URL, SUPABASE_KEY)

type Quiz struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"created_at"`
	Ques []Question `json:"question"`
	Msg  string
}

func (q *Quiz) Error() string {
	return q.Msg
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

func IndexHandler(c echo.Context) error {
	return c.Render(200, "index.html", "Hello, Qviz!")
}
