package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Count struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	Neutral  int `json:"neutral"`
}

func IncNegative(device string) {
	mu.Lock()
	if _, ok := counts[device]; !ok {
		counts[device] = Count{}
	}
	c := counts[device]
	c.Negative++
	counts[device] = c
	mu.Unlock()
}

func IncPositive(device string) {
	mu.Lock()
	if _, ok := counts[device]; !ok {
		counts[device] = Count{}
	}
	c := counts[device]
	c.Positive++
	counts[device] = c
	mu.Unlock()
}

func IncNeutral(device string) {
	mu.Lock()
	if _, ok := counts[device]; !ok {
		counts[device] = Count{}
	}
	c := counts[device]
	c.Neutral++
	counts[device] = c
	mu.Unlock()
}

func MessageHandler(deviceID string, msg int) {
	switch msg {
	case 1:
		IncNeutral(deviceID)
	case 2:
		IncPositive(deviceID)
	case 3:
		IncNegative(deviceID)
	}
}

var (
	mu     sync.Mutex
	counts = map[string]Count{}
)

func count(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	payload, err := json.Marshal(counts)
	mu.Unlock()
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
}

func reset(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	counts = map[string]Count{}
	mu.Unlock()
	w.WriteHeader(http.StatusOK)
}

func home(w http.ResponseWriter, r *http.Request) {
}

func Run(path string) {
	http.HandleFunc("/count", count)
	log.Println(http.Dir(path))
	http.Handle("/", http.FileServer(http.Dir(path)))
	http.HandleFunc("/reset", reset)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
