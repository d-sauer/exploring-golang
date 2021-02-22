package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Initiate web server
func main() {
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:9100",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// Route declaration
func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler)
	return r
}

type RootResponse struct {
	Id      uuid.UUID `json:uuid`
	Message string    `json:message`
	Time    time.Time `json:time`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	msg := RootResponse{
		Id:      uuid.New(),
		Time:    time.Now(),
		Message: "test",
	}

	json.NewEncoder(w).Encode(msg)
}
