package main

import (
	"log"
	"os"
	"time"

	mqttgo "github.com/eclipse/paho.mqtt.golang"
	"github.com/matias-pan-globant/feedback-button/mqtt"
	"github.com/matias-pan-globant/feedback-button/server"
)

func main() {
	var (
		uri   = "tcp://127.0.0.1:1883"
		topic = "feedback"
	)
	go server.Run(os.Getenv("STATIC"))

	go mqtt.Listen(uri, topic)

	log.Println("connecting")
	client := mqtt.Connect("pub", uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		log.Println("publishing")
		token := client.Publish(topic, 0, false, t.String())
		switch v := token.(type) {
		case *mqttgo.PublishToken:
			log.Println(v.Error())
		}
	}
}
