package ops

import (
	"context"
	"time"

	"github.com/d-sauer/exploring-go/fileuploader/internal/api"
)

type OperationController struct {
}

// Health check
// (GET /health)
func (ops *OperationController) HealthCheck(ctx context.Context, request api.HealthCheckRequestObject) (api.HealthCheckResponseObject, error) {
	status := "up"
	response := api.HealthCheck200JSONResponse{Status: &status}

	return response, nil
}

// Ping check
// (GET /ping)
func (ops *OperationController) PingCheck(ctx context.Context, request api.PingCheckRequestObject) (api.PingCheckResponseObject, error) {
	tnu := time.Now().UTC()
	response := api.PingCheck200JSONResponse{Timestamp: &tnu}

	return response, nil
}
