package services

//go:generate go run github.com/golang/mock/mockgen -destination ../../mocks/healthservice.go -package mocks . HealthService

import (
	"github.com/d-sauer/exploring-go/hello-k8s/hello-service/domain"
)

type Healther interface {
	GetHealth() (*domain.HealthStatus, error)
}

type DefaultHealther struct{}

func (dh *DefaultHealther) Health() (*domain.HealthStatus, error) {
	h := &domain.HealthStatus{
		State: domain.Up,
	}
	return h, nil
}
