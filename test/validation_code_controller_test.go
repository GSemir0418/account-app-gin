package test

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/controller"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVCCreate(t *testing.T) {
	setUpTestCase(t)

	tc := controller.ValidationCodeController{}
	tc.RegisterRoutes(r.Group("/api"))

	w := httptest.NewRecorder()
	body := &api.CreateValidationCodeRequest{
		Email: "1",
	}
	bodyJson, _ := json.Marshal(body)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/validation-codes",
		strings.NewReader(string(bodyJson)),
	)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
