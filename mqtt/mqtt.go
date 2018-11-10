package mqtt

import (
	"fmt"
	"log"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const (
	clientID = "broker"
)

type empty struct{}

var (
	devices = map[string]empty{}
)

// HandlerFunc is the function that handles
// new messages from topic of deviceID.
type HandlerFunc func(deviceID string, msg int)

// Client wraps an mqtt client.
type Client struct {
	client  mqtt.Client
	handler HandlerFunc
}

// New returns a new mqtt client.
func New(uri string) *Client {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetPassword("test")
	opts.SetUsername("test")
	opts.SetClientID(clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return &Client{
		client:  client,
		handler: logHandler,
	}
}

// Listen listens for new devices.
func (c *Client) Listen(topic string) {
	token := c.client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		deviceID := string(msg.Payload())
		if _, ok := devices[deviceID]; ok {
			return
		}
		c.handleDevice(deviceID)
	})
	log.Println(token.Error())
}

// HandleFunc sets the handler function.
func (c *Client) HandleFunc(f HandlerFunc) {
	c.handler = f
}

func (c *Client) handleDevice(deviceID string) {
	token := c.client.Subscribe(parseTopic(deviceID), 0, func(client mqtt.Client, msg mqtt.Message) {
		n, err := strconv.Atoi(string(msg.Payload()))
		if err != nil {
			log.Printf("Unable to parse msg to int. Got: %s", string(msg.Payload()))
		}
		c.handler(deviceID, n)
	})
	log.Println(token.Error())
}

func parseTopic(deviceID string) string {
	return fmt.Sprintf("%s/button", deviceID)
}

func logHandler(deviceID string, msg int) {
	log.Printf("Message for %s: %d", deviceID, msg)
}
