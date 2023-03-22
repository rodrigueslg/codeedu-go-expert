package main

import (
	"log"

	"github.com/rodrigueslg/codedu-goexpert/desafio-client-server-api/client"
	"github.com/rodrigueslg/codedu-goexpert/desafio-client-server-api/server"
)

func main() {
	initServer()
	client.RunClient()
}

func initServer() {
	repo := server.NewSQLiteRepository()
	service := server.NewService(repo)
	handler := server.NewCoinHandler(service)
	server := server.NewServer(handler)

	go func() {
		log.Fatal(server.Run())
	}()
}
