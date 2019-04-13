package server

import "log"

type queue struct {
	Clients chan *client
}

func NewQueue(size int) *queue {
	return &queue{
		Clients: make(chan *client, size),
	}
}

func (queue *queue) Push(element *client) {
	select {
	case queue.Clients <- element:
		go log.Println("Put client into queue. ID: " + element.ID.String())
	default:
		panic("Queue full")
	}
}

func (queue *queue) Pop() *client {
	select {
	case e := <-queue.Clients:
		go log.Println("Popped client from queue. ID: " + e.ID.String())
		return e
	default:
		panic("Queue empty")
	}
}
