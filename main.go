package main

import (
	"log"

	"bld/qviz/routes"
)

func main() {

	log.Println("Mainline logging")

	server := routes.NewServer()

	server.Logger.Fatal(server.Start(":4000"))

}
