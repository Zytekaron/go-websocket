package types

import (
	"encoding/json"
	"io"
)

type ServerMessage struct {
	Message *Message
	Client  *Client
}

type Message struct {
	Message string          `json:"message,omitempty"`
	Event   string          `json:"event,omitempty"`
	Code    int64           `json:"code,omitempty"`
	Data    json.RawMessage `json:"data,omitempty"`
}

type OutMessage struct {
	Message string      `json:"message,omitempty"`
	Event   string      `json:"event,omitempty"`
	Code    int64       `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ParseMessage(b []byte) (msg *Message, err error) {
	return msg, json.Unmarshal(b, &msg)
}

func ParseMessageReader(r io.Reader) (msg *Message, err error) {
	return msg, json.NewDecoder(r).Decode(&msg)
}
