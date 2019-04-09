package main

import (
	"log"

	"github.com/Nerzal/CardsOfBinokee-Server/pkg/server"
)

func main() {
	log.Println("Starting Server")

	server := server.NewServer(1337)

	server.Serve()

}
