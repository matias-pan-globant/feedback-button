package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.corp.globant.com/feedback-button/model"
)

var Count *model.Count = &model.Count{}

func GetCount(w http.ResponseWriter, r *http.Request) {
	payload, err := json.Marshal(Count)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(payload))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func StartHttpServer() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/count", GetCount).Methods(http.MethodGet)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
