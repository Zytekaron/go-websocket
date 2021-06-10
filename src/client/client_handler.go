package client

import (
	"fmt"
	"go-websocket/src/types"
)

// Handle incoming server messages
func makeHandler(server *types.Server) chan *types.Message {
	ch := make(chan *types.Message)
	go handle(server, ch)
	return ch
}

func handle(server *types.Server, ch chan *types.Message) {
	for msg := range ch {
		switch msg.Event {
		case "INIT":
			fmt.Println("C Init:", msg.Code, "|", msg.Message, "|", string(msg.Data))
		case "TEST":
			fmt.Println("C Test:", msg.Code, "|", msg.Message, "|", string(msg.Data))
			server.Send(&types.OutMessage{
				Message: "GOT",
				Event:   "TEST_REPLY",
			})
		case "TEST_REPLY":
			fmt.Println("C Test Reply:", msg.Code, "|", msg.Message, "|", string(msg.Data))
		case "":
			fmt.Println("C No event:", msg.Code, "|", msg.Message, "|", string(msg.Data))
		default:
			fmt.Println("C Unknown event:", msg.Code, "|", msg.Message, "|", string(msg.Data))
		}
	}
}
