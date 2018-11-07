package main

import (
	"github.corp.globant.com/feedback-button/http"
	"github.corp.globant.com/feedback-button/mqtt"
	"log"
	"net/url"
	"os"
	"time"
)

func main() {
	//uri, err := url.Parse(os.Getenv("MQTT_URL"))
	uri = "mqtt://<user>:<pass>@<server>.cloudmqtt.com:<port>/<topic>"
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "test"
	}

	go mqtt.Listen(uri, topic)

	client := mqtt.Connect("pub", uri)

	go http.StartHttpServer()

	// TODO remove this
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, t.String())
	}
}