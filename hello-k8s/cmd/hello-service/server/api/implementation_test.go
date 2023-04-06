package api

import (
	"github.com/d-sauer/exploring-go/hello-k8s/hello-service/domain"
	"github.com/d-sauer/exploring-go/hello-k8s/hello-service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	validator "openapi.tanna.dev/go/validator/openapi3"
	"testing"
)

func Test_healthRequestApi_GetHealth(t *testing.T) {
	mockService := mocks.NewMockHealthService(gomock.NewController(t))
	server := NewHealthRequestApi(mockService)

	t.Run("return UP", func(t *testing.T) {
		mockReturn := domain.HealthStatus{
			State: domain.Up,
		}
		mockService.EXPECT().GetHealth().Return(&mockReturn, nil)

		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/health", nil)

		server.GetHealth(rr, req, GetHealthParams{})

		t.Run("it returns 200", func(t *testing.T) {
			assert.Equal(t, 200, rr.Result().StatusCode)
		})

		t.Run("it maches OpenAPI", func(t *testing.T) {
			doc, err := GetSwagger()
			assert.NoError(t, err)

			_ = validator.NewValidator(doc).ForTest(t, rr, req)
		})
	})
}
