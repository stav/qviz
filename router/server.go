package router

import (
	"github.com/labstack/echo/v4"

	"bld/qviz/middleware"
)

// Exported function that returns an echo instance
func NewServer() *echo.Echo {
	// Create a new echo instance
	e := echo.New()
	e.Logger.Print("Logging to the echo logger")
	e.Renderer = newTemplate()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri} ${error}\n",
	}))

	// Static files
	e.File("/htmx.png", "images/htmx.png")
	e.File("/htmx.js", "js/htmx.js", middleware.AddScriptHeader)
	e.File("/head.js", "js/head.js", middleware.AddScriptHeader)

	// Public routes
	e.GET("/", IndexHandler)
	e.GET("/login", GetLoginHandler)
	e.POST("/login", PostLoginHandler)
	e.POST("/logout", PostLogoutHandler)
	e.GET("/register", GetRegisterHandler)
	e.POST("/register", PostRegisterHandler)

	// Guarded app routes
	app := e.Group("/app")
	app.Use(middleware.Sentry)
	app.GET("", AppIndexHandler)
	app.GET("/quiz/:quizId", AppQuizHandler)

	// Guarded API routes
	api := e.Group("/api")
	api.Use(middleware.Sentry)
	api.GET("/quiz/:quizId", ApiQuizHandler)
	api.GET("/quiz/:quizId/:questionNumber", ApiQuestionHandler)
	api.GET("/quiz/:quizId/:questionNumber/:answerId", ApiAnswerHandler)

	// Return the echo instance
	return e
}
