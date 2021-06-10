package server

import (
	"fmt"
	"go-websocket/src/types"
)

// fixme grab from config
var PoolSize = 1

// Handle incoming client messages
func makeHandler() chan *types.ServerMessage {
	ch := make(chan *types.ServerMessage)
	for i := 0; i < PoolSize; i++ {
		go handle(ch)
	}
	return ch
}

func handle(ch chan *types.ServerMessage) {
	for in := range ch {
		msg := in.Message
		c := in.Client

		switch msg.Event {
		case "INIT":
			fmt.Println("S Init:", msg.Code, "|", msg.Message, "|", string(msg.Data))
		case "TEST":
			fmt.Println("S Test:", msg.Code, "|", msg.Message, "|", string(msg.Data))
			c.Send(&types.OutMessage{
				Message: "GOT",
				Event:   "TEST_REPLY",
			})
		case "TEST_REPLY":
			fmt.Println("S Test Reply:", msg.Code, "|", msg.Message, "|", string(msg.Data))
		case "":
			fmt.Println("S No event:", msg.Code, "|", msg.Message, "|", string(msg.Data))
		default:
			fmt.Println("S Unknown event:", msg.Code, "|", msg.Message, "|", string(msg.Data))
		}
	}
}
