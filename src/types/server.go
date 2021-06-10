package types

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Server struct {
	Conn *websocket.Conn
	send chan *OutMessage
}

func NewServer(conn *websocket.Conn) *Server {
	c := &Server{
		Conn: conn,
		send: make(chan *OutMessage),
	}
	c.startHandlers()
	return c
}

func (s *Server) Send(msg *OutMessage) {
	go func() {
		s.send <- msg
	}()
}

func (s *Server) Stop() {
	close(s.send)
}

func (s *Server) startHandlers() {
	go func() {
		for msg := range s.send {
			err := s.Conn.WriteJSON(msg)
			if err != nil {
				fmt.Println("Server write pump send error:", err)
			}
		}
	}()
}
