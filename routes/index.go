package routes

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	supa "github.com/nedpals/supabase-go"

	_ "github.com/joho/godotenv/autoload"

	qviz_middleware "bld/qviz/middleware"
)

var SUPABASE_URL string = os.Getenv("SUPABASE_URL")
var SUPABASE_KEY string = os.Getenv("SUPABASE_KEY")
var supabase = supa.CreateClient(SUPABASE_URL, SUPABASE_KEY)

type Quiz struct {
	ID   int        `json:"id"`
	Name string     `json:"name"`
	Date string     `json:"created_at"`
	Ques []Question `json:"question"`
	Msg  string
}

func (q *Quiz) Error() string {
	return q.Msg
}

type Question struct {
	ID   int      `json:"id"`
	Text string   `json:"text"`
	Ans  []Answer `json:"answer"`
}

type Answer struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	True bool   `json:"is_correct"`
}

func IndexHandler(c echo.Context) error {
	return c.Render(200, "index.html", "Hello, Qviz!")
}

// export a function that returns an echo instance
func NewServer() *echo.Echo {
	e := echo.New()
	e.Logger.Print("Logging to the echo logger")
	e.Renderer = newTemplate()
	// e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri} ${error}\n",
	}))

	e.File("/htmx.png", "images/htmx.png")
	e.File("/htmx.js", "js/htmx.js", qviz_middleware.AddScriptHeader)
	e.File("/head.js", "js/head.js", qviz_middleware.AddScriptHeader)

	e.GET("/", GetLoginHandler)
	e.GET("/login", GetLoginHandler)
	e.POST("/login", PostLoginHandler)
	e.POST("/logout", PostLogoutHandler)
	e.GET("/register", GetRegisterHandler)
	e.POST("/register", PostRegisterHandler)

	app := e.Group("/app")
	app.Use(qviz_middleware.Sentry)
	app.GET("", AppIndexHandler)
	app.GET("/quiz/:id", AppQuizHandler)

	api := e.Group("/api")
	api.Use(qviz_middleware.Sentry)
	api.GET("/quiz/:id", ApiQuizHandler)

	// return the echo instance
	return e
}
