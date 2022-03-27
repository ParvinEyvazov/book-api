package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/parvineyvazov/book-api/server"
)

func main() {

	server := server.NewServer()
	log.Fatalln(server.StartServer(8080))
}
