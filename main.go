package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"

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

func main() {

	log.Println("Mainline logging")

	e := echo.New()
	e.Logger.Print("Logging to the echo logger")
	e.Renderer = newTemplate()

	// e.Use(middleware.Recover())
	// e.Use(echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte("secret"),
	// 	ContinueOnIgnoredError: true,
	// 	SuccessHandler: func(c echo.Context) {
	// 		fmt.Println("SuccessHandler")
	// 		user := c.Get("user").(*jwt.Token)
	// 		claims := user.Claims.(jwt.MapClaims)
	// 		// c.Logger().Info("User: ", claims["name"])
	// 		log.Println("user:::", claims["name"])
	// 	},
	// 	ErrorHandler: func(c echo.Context, err error) error {
	// 		fmt.Println("ErrorHandler")
	// 		c.JSON(http.StatusUnauthorized, map[string]string{
	// 			"error": err.Error(),
	// 		})
	// 		return nil
	// 	},
	// }))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri} ${error}\n",
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the credentials from HTTP request header and perform a security
			// check
			hxrequest := c.Request().Header.Get("Hx-Request")
			fmt.Println("hxrequest::", hxrequest)

			// For invalid credentials
			// return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

			// For valid credentials call next
			return next(c)
		}
	})

	e.File("/htmx.png", "images/htmx.png")
	e.File("/htmx.js", "js/htmx.js")

	e.GET("/", routes.IndexHandler)
	e.GET("/login", routes.GetLoginHandler)
	e.POST("/login", routes.PostLoginHandler)
	e.GET("/register", routes.GetRegisterHandler)
	e.POST("/register", routes.PostRegisterHandler)

	e.GET("/jwt", func(c echo.Context) error {
		var user = c.Get("user")
		fmt.Println("user::", user)
		token, ok := user.(*jwt.Token) // by default token is stored under `user` key
		if !ok {
			return errors.New("JWT token missing or invalid")
		}
		claims, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
		if !ok {
			return errors.New("failed to cast claims as jwt.MapClaims")
		}
		return c.JSON(http.StatusOK, claims)
	})

	e.Logger.Fatal(e.Start(":8888"))

}
