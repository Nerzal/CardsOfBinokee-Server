package server

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Server interface {
	Serve() error
}

type server struct {
	tlsConfig      *tls.Config
	queue          *queue
	activeSessions []*session
}

// NewServer constructs a new instance of server which fullfills the requirements of the Server interface
func NewServer(tlsConfig *tls.Config) Server {
	queue := NewQueue(100)
	return &server{
		queue:     queue,
		tlsConfig: tlsConfig,
	}
}

func (server *server) Serve() error {
	l, err := tls.Listen("tcp", ":995", server.tlsConfig)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer l.Close()

	go server.checkQueue()

	for {
		newConnection, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go server.handleNewConnection(newConnection)
	}
}

func (server *server) handleNewConnection(newConnection net.Conn) {
	newID := uuid.Must(uuid.NewV4())

	newClient := &client{
		Connection: newConnection,
		ID:         newID,
	}

	server.queue.Push(newClient)

	go readPackage(newConnection)
}

func (server *server) checkQueue() {
	for {
		// log.Println("Queue: current queue length: " + strconv.Itoa(len(server.queue.Clients)))
		if len(server.queue.Clients) >= 2 {
			newSession := createNewSession(server.queue)
			server.activeSessions = append(server.activeSessions, newSession)
			go log.Println("Created new Session with ID: " + newSession.ID.String())
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func createNewSession(queue *queue) *session {
	newID := uuid.Must(uuid.NewV4())
	player1 := queue.Pop()
	player2 := queue.Pop()

	return &session{
		ID:      &newID,
		Player1: player1,
		Player2: player2,
	}
}

func readPackage(c net.Conn) {
	buf := make([]byte, 512)
	n, err := c.Read(buf)
	if err != nil {
		go log.Printf("server: conn: read: %s", err)
	}

	go log.Printf("server: conn: echo %q\n", string(buf[:n]))

	_, err = c.Write([]byte("Hallo Client, du listiger Lurch! :) :-*"))
	if err != nil {
		go log.Println("Ohh noees!!!111elf I failed, sry bro!  ", err)
	}
	c.Close()
}
