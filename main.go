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
		topic = "up"
	)
	go server.Run(os.Getenv("STATIC"))

	client := mqtt.New(uri)
	client.HandleFunc(server.MessageHandler)
	client.Listen(topic)

	opts := mqttgo.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetPassword("test")
	opts.SetUsername("test")
	opts.SetClientID("wimo")
	pub := mqttgo.NewClient(opts)
	token := pub.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	pub.Publish(topic, 0, false, "device-1")
	ticker := time.NewTicker(time.Second * 2)
	for range ticker.C {
		pub.Publish("device-1/button", 0, false, "1")
		pub.Publish("device-1/button", 0, false, "2")
		pub.Publish("device-1/button", 0, false, "3")
	}
}
