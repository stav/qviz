package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bld/qviz/routes"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	return &Template{
		tmpl: template.Must(template.ParseGlob("views/*.html")),
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func printHeaders(header http.Header) {
	keys := make([]string, 0, len(header))
	for k := range header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%-20s: %#v\n", k, header.Get(k))
	}
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

	e.GET("/", routes.IndexHandler)
	e.GET("/login", routes.GetLoginHandler)
	e.POST("/login", routes.PostLoginHandler)
	e.GET("/register", routes.GetRegisterHandler)
	e.POST("/register", routes.PostRegisterHandler)

	app := e.Group("/app")
	app.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("\n____________________")
			fmt.Println("Middleware: Headers")
			fmt.Println()

			request := c.Request()
			headers := request.Header

			printHeaders(headers)

			cookie, err := c.Cookie("token")
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			// "token=eyJhbGciOiJI---.---" 877 chars
			token := cookie.Value
			fmt.Println("\nToken: ", token)

			// Now perform a security check on the token

			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})

	app.GET("", func(c echo.Context) error {
		fmt.Println("GET /app")
		return c.JSON(http.StatusOK, "{ \"message\": \"GET /app\" }")
	})

	e.Logger.Fatal(e.Start(":8888"))

}
