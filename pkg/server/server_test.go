package server

import (
	"io/ioutil"
	"log"
	"net"
	"testing"
)

func TestFoo(t *testing.T) {
	server := NewServer(1337)
	go server.Serve()

	ip, _, err := net.ParseCIDR("127.0.0.1/32")
	if err != nil {
		t.Error("Failed to parse IP", err)
	}

	clientConnection, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   ip,
		Port: 1337,
	})

	if err != nil {
		t.Error("Failed to parse IP", err)
	}

	clientRead(clientConnection)

}

func clientRead(connection *net.TCPConn) {
	for {
		messageBytes, err := ioutil.ReadAll(connection)
		if err != nil {
			log.Println("Failed to read from connection: ", err)
		} else {
			log.Println("Client: " + string(messageBytes))
			return
		}
	}
}
