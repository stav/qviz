package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	qviz_middleware "bld/qviz/middleware"
)

func IndexHandler(c echo.Context) error {
	return c.Render(200, "index.html", "Hello, Qviz!")
}

// Exported function that returns an echo instance
func NewServer() *echo.Echo {
	// Create a new echo instance
	e := echo.New()
	e.Logger.Print("Logging to the echo logger")
	e.Renderer = newTemplate()
	// e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri} ${error}\n",
	}))

	// Static files
	e.File("/htmx.png", "images/htmx.png")
	e.File("/htmx.js", "js/htmx.js", qviz_middleware.AddScriptHeader)
	e.File("/head.js", "js/head.js", qviz_middleware.AddScriptHeader)

	// Public routes
	e.GET("/", IndexHandler)
	e.GET("/login", GetLoginHandler)
	e.POST("/login", PostLoginHandler)
	e.POST("/logout", PostLogoutHandler)
	e.GET("/register", GetRegisterHandler)
	e.POST("/register", PostRegisterHandler)

	// Guarded app routes
	app := e.Group("/app")
	app.Use(qviz_middleware.Sentry)
	app.GET("", AppIndexHandler)
	app.GET("/quiz/:id", AppQuizHandler)

	// Guarded API routes
	api := e.Group("/api")
	api.Use(qviz_middleware.Sentry)
	api.GET("/quiz/:id", ApiQuizHandler)

	// Return the echo instance
	return e
}
