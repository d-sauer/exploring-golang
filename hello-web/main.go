package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Initiate web server
func main() {
	const port = 2200

	log.Println("Starting server on port: ", port)
	router := router()
	srv := &http.Server{
		Handler:      router,
		Addr:         fmt.Sprint("127.0.0.1:", port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// Route declaration
func router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", rootHandler)

	router.Use(prometheusMiddleware)
	router.Path("/metrics").Handler((promhttp.Handler()))
	// r.HandleFunc("/metrics", promhttp.Handler())

	// r.Walk(play)
	return router
}

// func play(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
// 	log.Println("Route name: ", route.GetName(), "   route:", route, "  router:", router, "  ancestors:", ancestors)
// 	return nil
// }

type RootResponse struct {
	Id      uuid.UUID `json:uuid`
	Message string    `json:message`
	Time    time.Time `json:time`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")

	w.Header().Set("Content-Type", "application/json")

	msg := RootResponse{
		Id:      uuid.New(),
		Time:    time.Now(),
		Message: "test",
	}

	json.NewEncoder(w).Encode(msg)
}
