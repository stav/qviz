package main

import (
	"log"

	"bld/qviz/routes"
)

func main() {

	log.Println("Mainline logging")

	e := routes.NewServer()

	e.Logger.Fatal(e.Start(":4000"))

}
