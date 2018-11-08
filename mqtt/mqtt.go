package mqtt

import (
	"log"
	"math/rand"
	"time"

	"github.com/eclipse/paho.mqtt.golang"
	"github.com/matias-pan-globant/feedback-button/server"
)

func Connect(clientId string, uri string) mqtt.Client {
	opts := clientOpts(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Println("something bad happened")
		log.Fatal(err)
	}
	return client
}

func clientOpts(clientId string, uri string) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetPassword("test")
	opts.SetUsername("test")
	opts.SetClientID(clientId)
	return opts
}

func Listen(uri string, topic string) {
	client := Connect("sub", uri)
	token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		switch rand.Intn(3) {
		case 0:
			server.IncPositive("device-1")
		case 1:
			server.IncNegative("device-1")
		case 2:
			server.IncNeutral("device-1")
		}
	})
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Println("something bad happened with listen")
		log.Fatal(err)
	}
	log.Println("leaving listen")
}
