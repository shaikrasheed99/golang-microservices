package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shaikrasheed99/broker-service/app/models/responses"
	"github.com/shaikrasheed99/broker-service/utils/constants"
	"github.com/stretchr/testify/assert"
)

func TestBrokerHandler_Health(t *testing.T) {
	t.Run("should be able to up and running", func(t *testing.T) {
		handler := NewBrokerHandler()

		router := gin.Default()
		router.GET(constants.HealthEndpoint, handler.Health)

		req, _ := http.NewRequest("GET", constants.HealthEndpoint, nil)
		res := httptest.NewRecorder()
		router.ServeHTTP(res, req)

		var resBody responses.SuccessResponse
		_ = json.Unmarshal(res.Body.Bytes(), &resBody)

		assert.Equal(t, constants.Success, resBody.Status)
		assert.Equal(t, http.StatusText(http.StatusOK), resBody.Code)
		assert.Equal(t, constants.UP, resBody.Message)
		assert.Equal(t, nil, resBody.Data)
	})
}