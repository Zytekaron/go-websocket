package types

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	send chan *OutMessage
}

func NewClient(id string, conn *websocket.Conn) *Client {
	c := &Client{
		ID:   id,
		Conn: conn,
		send: make(chan *OutMessage),
	}
	c.startHandlers()
	return c
}

func (c *Client) Send(msg *OutMessage) {
	go func() {
		c.send <- msg
	}()
}

func (c *Client) Stop() {
	close(c.send)
}

func (c *Client) startHandlers() {
	go func() {
		for msg := range c.send {
			err := c.Conn.WriteJSON(msg)
			if err != nil {
				fmt.Println("Client write pump send error:", err)
			}
		}
	}()
}
