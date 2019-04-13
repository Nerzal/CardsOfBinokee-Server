package server

import (
	"net"

	uuid "github.com/satori/go.uuid"
)

type client struct {
	ID         uuid.UUID
	Connection net.Conn
}
