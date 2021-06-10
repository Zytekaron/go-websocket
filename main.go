package main

import (
	"go-websocket/src/client"
	"go-websocket/src/server"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	go server.Run()

	<-time.After(time.Millisecond)

	go client.Run(1)
	go client.Run(2)

	select {}
}
