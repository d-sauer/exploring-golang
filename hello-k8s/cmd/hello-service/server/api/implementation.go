package api

import (
	"encoding/json"
	"github.com/d-sauer/exploring-go/hello-k8s/hello-service/domain/services"
	"net/http"
	"time"
)

type healthRequestApi struct {
	service services.Healther
}

func NewHealthRequestApi(service services.Healther) ServerInterface {
	return &healthRequestApi{
		service: service,
	}
}

func (h *healthRequestApi) GetHealth(w http.ResponseWriter, r *http.Request) {
	status, err := h.service.GetHealth()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if status == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	time := time.Now()
	response := Health{
		Status:    string(status.State),
		Timestamp: &time,
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(bytes)
	w.WriteHeader(http.StatusOK)
}
