package main

import (
	"github.com/matias-pan-globant/feedback-button/mqtt"
	"github.com/matias-pan-globant/feedback-button/server"
)

func main() {
	var (
		uri   = "tcp://127.0.0.1:1883"
		topic = "up"
	)

	client := mqtt.New(uri)
	client.HandleFunc(server.MessageHandler)
	client.Listen(topic)

	server.Run("./server/static")
}
