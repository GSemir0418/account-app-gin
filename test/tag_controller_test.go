package test

import (
	"account-app-gin/internal/controller"
	"account-app-gin/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTagCreate(t *testing.T) {
	setUpTestCase(t)
	// 注册路由
	tc := controller.TagController{}
	tc.RegisterRoutes(r.Group("/api"))
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
	tagJson, _ := json.Marshal(tag)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/tags",
		strings.NewReader(string(tagJson)),
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

func TestTagUpdate(t *testing.T) {
	setUpTestCase(t)
	// 注册路由
	tc := controller.TagController{}
	tc.RegisterRoutes(r.Group("/api"))
	// 初始化 w
	w := httptest.NewRecorder()
	// 创建一个 user
	user := &database.User{
		Email: "1@qq.com",
	}
	if tx := database.DB.Create(user); tx.Error != nil {
		t.Fatal("Create user failed:", tx.Error)
	}
	// 创建一个 tag
	tag := &database.Tag{
		UserID: user.ID,
		Sign:   "⌚️",
		Name:   "电子产品",
		Kind:   "expenses",
	}
	tx := database.DB.Create(tag)
	if tx.Error != nil {
		t.Fatal("Create tag failed:", tx.Error)
	}

	// 更新后的 tag
	newTag := map[string]string{
		"sign": "🏮",
		"name": "新名称",
	}

	newTagJson, _ := json.Marshal(newTag)
	req, _ := http.NewRequest(
		"PATCH",
		fmt.Sprintf("/api/v1/tags/%d", tag.ID),
		strings.NewReader(string(newTagJson)),
	)
	// 发起请求
	r.ServeHTTP(w, req)
	// 处理响应体
	var response map[string]interface{}
	assert.Equal(t, 200, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("json.Unmarshal fail %v", err)
	}
	assert.Equal(t, response["UserID"], float64(user.ID))
	assert.Equal(t, response["Sign"], newTag["sign"])
	assert.Equal(t, response["Kind"], tag.Kind)
}

// func TestItemPaged(t *testing.T) {
// 	setUpTestCase(t)
// 	// 注册路由
// 	ic := controller.ItemController{}
// 	ic.RegisterRoutes(r.Group("/api"))
// 	// 初始化 w
// 	w := httptest.NewRecorder()
// 	// 构造请求
// 	req, _ := http.NewRequest(
// 		"GET",
// 		"/api/v1/items?page=3&page_size=5",
// 		nil,
// 	)
// 	// 创建一个 user
// 	user := &database.User{
// 		Email: "1@qq.com",
// 	}
// 	tx := database.DB.Create(user)
// 	if tx.Error != nil {
// 		t.Fatal("Create user failed:", tx.Error)
// 	}
// 	// 创建一个 tag
// 	tag := &database.Tag{
// 		UserID: user.ID,
// 		Sign:   "⌚️",
// 		Name:   "电子产品",
// 		Kind:   "expenses",
// 	}
// 	tx = database.DB.Create(tag)
// 	if tx.Error != nil {
// 		t.Fatal("Create tag failed:", tx.Error)
// 	}

// 	// 创建 13 个 item
// 	for i := 0; i < int(13); i++ {
// 		item := &database.Item{
// 			UserID:     user.ID,
// 			Amount:     10000,
// 			Kind:       "expenses",
// 			Tags:       []*database.Tag{tag},
// 			HappenedAt: time.Now(),
// 		}
// 		if tx = database.DB.Create(item); tx.Error != nil {
// 			t.Fatal(tx.Error)
// 		}
// 	}

// 	// 发起请求
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, 200, w.Code)

// 	var response api.GetPagedResponse
// 	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
// 		t.Fatalf("json.Unmarshal fail %v", err)
// 	}
// 	// 先用断言 后面会补充类型
// 	assert.Equal(t, 3, len(response.Resources))
// }
