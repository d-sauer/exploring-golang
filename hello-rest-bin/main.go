package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gorilla/mux"
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
	router.HandleFunc("/time", timeHandler)
	router.HandleFunc("/dump", dumpHandler)
	router.HandleFunc("/dump/{rest:.*}", dumpHandler)

	return router
}

type TimeResponse struct {
	Time time.Time `json:time`
}

func dumpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URI: %s  >> Remaining path: %s \n", r.URL.Path, mux.Vars(r)["rest"])
	dumped, err := httputil.DumpRequest(r, true)

	bodyB, _ := ioutil.ReadAll(r.Body)
	bodyStr := string(bytes.Replace(bodyB, []byte("\r"), []byte("\r\n"), -1))
	fmt.Println(bodyStr)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("can map request data"))
		return
	}
	w.Write(dumped)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received request")

	w.Header().Set("Content-Type", "application/json")

	msg := TimeResponse{
		Time: time.Now(),
	}
	json.NewEncoder(w).Encode(msg)
}
