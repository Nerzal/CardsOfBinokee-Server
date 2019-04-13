package server

import (
	"github.com/satori/go.uuid"
)

type session struct {
	ID        *uuid.UUID
	Player1   *client
	Player2   *client
	Specators []client
}
