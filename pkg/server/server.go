package server

import (
	"log"
	"net"
	"strconv"
)

type Server interface {
	Serve()
}

type server struct {
	port int
}

func NewServer(port int) Server {
	return &server{
		port: port,
	}
}

func (server *server) Serve() {
	listener, _ := net.Listen("tcp", ":"+strconv.Itoa(server.port))
	log.Println("Listening for new Connections on Port: " + strconv.Itoa(server.port))
	for {
		newConnection, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection: ", err.Error())
			// return err
		}

		newClient := NewClient(newConnection)
		newClient.Listen()

		log.Println("Server: New client has connected!")
		
		_, err = newConnection.Write([]byte("Deine muddah!"))
		if err != nil {
			log.Println("Failed to write message to client: ", err)
		}
	}
}
