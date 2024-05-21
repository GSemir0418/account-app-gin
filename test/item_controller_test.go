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
	// 创建一个 user
	user := &database.User{
		Email: "1@qq.com",
	}
	tx := database.DB.Create(user)
	if tx.Error != nil {
		t.Fatal("Create user failed:", tx.Error)
	}
	// 创建一个 tag
	tag := &database.Tag{
		UserID: user.ID,
		Sign:   "⌚️",
		Name:   "电子产品",
		Kind:   "expenses",
	}
	tx = database.DB.Create(tag)
	if tx.Error != nil {
		t.Fatal("Create tag failed:", tx.Error)
	}

	// 创建 item
	item := &database.Item{
		Amount:     100,
		Tags:       []*database.Tag{tag},
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

func TestItemPaged(t *testing.T) {
	setUpTestCase(t)
	// 注册路由
	ic := controller.ItemController{}
	ic.RegisterRoutes(r.Group("/api"))
	// 初始化 w
	w := httptest.NewRecorder()
	// 构造请求
	req, _ := http.NewRequest(
		"GET",
		"/api/v1/items?page=3&page_size=5",
		nil,
	)
	// 创建一个 user
	user := &database.User{
		Email: "1@qq.com",
	}
	tx := database.DB.Create(user)
	if tx.Error != nil {
		t.Fatal("Create user failed:", tx.Error)
	}
	// 创建一个 tag
	tag := &database.Tag{
		UserID: user.ID,
		Sign:   "⌚️",
		Name:   "电子产品",
		Kind:   "expenses",
	}
	tx = database.DB.Create(tag)
	if tx.Error != nil {
		t.Fatal("Create tag failed:", tx.Error)
	}

	// 创建 13 个 item
	for i := 0; i < int(13); i++ {
		item := &database.Item{
			UserID:     user.ID,
			Amount:     10000,
			Kind:       "expenses",
			Tags:       []*database.Tag{tag},
			HappenedAt: time.Now(),
		}
		if tx = database.DB.Create(item); tx.Error != nil {
			t.Fatal(tx.Error)
		}
	}

	// 发起请求
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var response api.GetPagedResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("json.Unmarshal fail %v", err)
	}
	// 先用断言 后面会补充类型
	assert.Equal(t, 3, len(response.Resources))
}
