package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go-websocket/src/types"
	"net/http"
	"strconv"
)

// Run a client that connects to the server
func Run(which int) {
	id := "#" + strconv.Itoa(which)

	uri := "ws://127.0.0.1:1337/ws"
	conn, _, err := websocket.DefaultDialer.Dial(uri, http.Header{
		"Authorization": []string{"Bearer i_am_god"},
	})
	if err != nil {
		fmt.Println("Dial error:", err)
		return
	}
	defer conn.Close()

	server := types.NewServer(conn)
	ch := makeHandler(server)

	server.Send(&types.OutMessage{
		Event:   "INIT",
		Message: "Initialization from Client " + id,
		Data:    []interface{}{"This is a test!", 123},
	})

	for {
		var msg *types.Message
		err = conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println("C "+id+" Read error:", err)
			break
		}

		ch <- msg
	}
}
