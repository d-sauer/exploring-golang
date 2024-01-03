package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/d-sauer/exploring-go/fileuploader/internal/api"
	"github.com/d-sauer/exploring-go/fileuploader/internal/ops"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime/middleware"
	chimiddleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	fmt.Println("Hello..")
	newHttpServer()
}

func newHttpServer() {
	ops := NewOperationController()
	server := NewServer(ops)
	router := chi.NewRouter()

	// Add swagger UI endpoints
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(swagger)
	})
	router.Handle("/swagger/", middleware.SwaggerUI(middleware.SwaggerUIOpts{
		Path:    "/swagger/",
		SpecURL: "/swagger/doc.json",
	}, nil))

	// Enable validation of incoming requests
	validator := chimiddleware.OapiRequestValidatorWithOptions(
		swagger,
		&chimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
					return nil
				},
			},
		},
	)

	apiServer := api.HandlerWithOptions(
		api.NewStrictHandler(server, nil),
		api.ChiServerOptions{
			BaseURL:    "/api", // must match openapi spec server url
			BaseRouter: router,
			Middlewares: []api.MiddlewareFunc{
				validator,
			},
		},
	)

	addr := ":8000"
	httpServer := http.Server{
		Addr:    addr,
		Handler: apiServer,
	}

	log.Println("Server listening on", addr)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}

func NewServer(operationController *ops.OperationController) api.StrictServerInterface {
	return &server{
		OperationController: operationController,
	}
}

func NewOperationController() *ops.OperationController {
	return &ops.OperationController{}
}

type server struct {
	*ops.OperationController
}
