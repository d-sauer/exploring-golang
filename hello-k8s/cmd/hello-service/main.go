package main

import (
	"flag"
	"fmt"
	"github.com/d-sauer/exploring-go/hello-k8s/hello-service/domain/services"
	"github.com/d-sauer/exploring-go/hello-k8s/hello-service/server/api"
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

//
//func main() {
//	var port string
//	flag.StringVar(&port, "port", "9070", "Define custom server port")
//	flag.Parse()
//
//	log.Print("Starting server on port:" + port)
//	r := mux.NewRouter()
//	h := server.NewHandler()
//	r.Handle("/app", h).Name("application")
//
//	srv := &http.Server{
//		Handler:      r,
//		Addr:         "127.0.0.1:" + port,
//		WriteTimeout: 15 * time.Second,
//		ReadTimeout:  15 * time.Second,
//	}
//	log.Fatal(srv.ListenAndServe())
//	log.Print("Server stopped")
//}

func main() {
	var port = flag.Int("port", 9070, "Port for test HTTP server")
	flag.Parse()

	openapi, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading OpenAPI spec\n: %s", err)
		os.Exit(1)
	}

	openapi.Servers = nil

	var service services.Healther

	healthRequestApiService := api.NewHealthRequestApi(service)

	r := mux.NewRouter()

	// enforce consumers use the right contract, required on top of checks that are done in the generated API code
	r.Use(middleware.OapiRequestValidator(openapi))

	api.HandlerFromMux(healthRequestApiService, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	log.Fatal(s.ListenAndServe())
}
