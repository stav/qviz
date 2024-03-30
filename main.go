package main

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bld/qviz/routes"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func main() {

	log.Println("Mainline logging")

	e := echo.New()
	e.Logger.Print("Logging to the echo logger")
	e.Renderer = newTemplate()

	// e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri} ${error}\n",
	}))

	e.File("/htmx.png", "images/htmx.png")
	e.File("/htmx.js", "js/htmx.js")

	e.GET("/client.js", routes.ScriptHandler)

	e.GET("/", routes.IndexHandler)
	e.GET("/login", routes.GetLoginHandler)
	e.POST("/login", routes.PostLoginHandler)
	e.GET("/register", routes.GetRegisterHandler)
	e.POST("/register", routes.PostRegisterHandler)

	e.Logger.Fatal(e.Start(":8888"))

}
