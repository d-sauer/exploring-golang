package server

import (
	"fmt"
	"github.com/d-sauer/exploring-go/hello-k8s/hello-service/domain/services"
	"github.com/d-sauer/exploring-go/hello-k8s/hello-service/server/api"
	"net/http"
	"os"

	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/gorilla/mux"
)

func NewHandler() http.Handler {
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

	return api.HandlerFromMux(healthRequestApiService, r)
}
