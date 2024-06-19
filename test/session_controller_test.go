package test

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/controller"
	"account-app-gin/internal/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSession(t *testing.T) {
	setUpTestCase(t)

	sc := controller.SessionController{}
	sc.RegisterRoutes(r.Group("/api"))

	// 生成一个验证码
	vcode := database.ValidationCode{
		Email: "845217811@qq.com",
		Code:  "1234",
	}
	database.DB.Create(&vcode)

	w := httptest.NewRecorder()
	body := &api.SessionRequest{
		Email: "845217811@qq.com",
		Code:  "1234",
	}
	bodyJson, _ := json.Marshal(body)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/session",
		strings.NewReader(string(bodyJson)),
	)
	r.ServeHTTP(w, req)

	var res api.SessionResponse
	json.Unmarshal(w.Body.Bytes(), &res)

	assert.Equal(t, 200, w.Code)
	// 断言响应体存在 jwt 且类型为字符串
	assert.NotEmpty(t, res.Jwt)
	assert.NotEmpty(t, res.UserID)
	assert.IsType(t, "", res.Jwt)
}
