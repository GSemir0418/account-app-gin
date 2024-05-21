package test

import (
	"account-app-gin/internal/controller"
	"account-app-gin/internal/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestItemCreate(t *testing.T) {
	setUpTestCase(t)
	// 注册路由
	ic := controller.ItemController{}
	ic.RegisterRoutes(r.Group("/api"))
	// 初始化 w
	w := httptest.NewRecorder()
	// 创建一个 user, 默认一定成功
	user := &database.User{
		Email: "1@qq.com",
	}
	tx := database.DB.Create(user)
	if tx.Error != nil {
		t.Fatal("Create user failed:", tx.Error)
	}

	// 创建 item
	item := &database.Item{
		Amount:     100,
		TagIDs:     "1,2,3",
		UserID:     user.ID,
		Kind:       "in_come",
		HappenedAt: time.Now(),
	}
	itemJson, _ := json.Marshal(item)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/items",
		strings.NewReader(string(itemJson)),
	)
	// 发起请求
	r.ServeHTTP(w, req)
	// 处理响应体
	var response map[string]interface{}
	assert.Equal(t, 200, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, response["UserID"], float64(user.ID))
}
