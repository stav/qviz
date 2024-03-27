package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"bld/qviz/routes"
)

func main() {

	log.Println("Mainline logging")

	e := echo.New()
	e.Logger.Print("Logging to the echo logger")

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} ${method} ${uri} ${error}\n",
	}))

	e.GET("/", routes.IndexHandler)

	e.Logger.Fatal(e.Start(":8888"))

}
