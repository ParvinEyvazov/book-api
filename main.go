package main

import (
	"log"

	"github.com/parvineyvazov/book-api/server"
)

func main() {

	server := server.NewServer()
	log.Fatalln(server.StartServer(8081))
}
