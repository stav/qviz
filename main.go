package main

import (
	"log"

	"bld/qviz/router"
)

func main() {

	log.Println("Mainline logging")

	server := router.NewServer()

	server.Logger.Fatal(server.Start(":4000"))

}
