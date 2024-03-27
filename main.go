package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	log.Println("Mainline logging")

	e := echo.New()
	e.Logger.Print("Logging to the echo logger")

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri} ${error}\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8888"))

}
