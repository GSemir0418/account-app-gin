package test

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/controller"
	"account-app-gin/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

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
	body := &api.CreateTagRequest{
		UserID: user.ID,
		Sign:   "⌚️",
		Name:   "电子产品",
		Kind:   "expenses",
	}
	bodyJson, _ := json.Marshal(body)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/tags",
		strings.NewReader(string(bodyJson)),
	)
	// 发起请求
	r.ServeHTTP(w, req)
	// 处理响应体
	var response map[string]interface{}
	assert.Equal(t, 200, w.Code)
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, response["userId"], float64(user.ID))
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
	body := api.UpdateTagRequest{
		Sign: strPtr("🏮"),
		Name: strPtr("新名称"),
	}

	bodyJson, _ := json.Marshal(body)
	req, _ := http.NewRequest(
		"PATCH",
		fmt.Sprintf("/api/v1/tags/%d", tag.ID),
		strings.NewReader(string(bodyJson)),
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
	assert.Equal(t, response["userId"], float64(user.ID))
	assert.Equal(t, response["sign"], *body.Sign)
	assert.Equal(t, response["name"], *body.Name)
	assert.Equal(t, response["kind"], tag.Kind)
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

func TestDeleteTag(t *testing.T) {
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
	tx = database.DB.Create(tag)
	if tx.Error != nil {
		t.Fatal("Create tag failed:", tx.Error)
	}
	// 创建一个 item
	item := &database.Item{
		Amount:     1000000,
		Tags:       []*database.Tag{tag},
		UserID:     user.ID,
		Kind:       "in_come",
		HappenedAt: time.Now(),
	}
	tx = database.DB.Create(item)
	if tx.Error != nil {
		t.Fatal("Create item failed:", tx.Error)
	}
	// 测试此时三个表分别只有 1 条数据
	var tagCount int64
	database.DB.Model(&database.Tag{}).Count(&tagCount)
	var itemCount int64
	database.DB.Model(&database.Item{}).Count(&itemCount)
	var itemTagsCount int64
	database.DB.Raw("SELECT COUNT(*) FROM item_tags").Scan(&itemTagsCount)
	assert.Equal(t, int64(1), tagCount)
	assert.Equal(t, int64(1), itemCount)
	assert.Equal(t, int64(1), itemTagsCount)
	// 构造请求
	req, _ := http.NewRequest(
		"DELETE",
		fmt.Sprintf("/api/v1/tags/%d", tag.ID),
		nil,
	)
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	database.DB.Model(&database.Tag{}).Count(&tagCount)
	database.DB.Model(&database.Item{}).Count(&itemCount)
	database.DB.Raw("SELECT COUNT(*) FROM item_tags").Scan(&itemTagsCount)
	// 断言删除成功，并且关联表的数据也会删除
	assert.Equal(t, int64(0), tagCount)
	assert.Equal(t, int64(1), itemCount)
	assert.Equal(t, int64(0), itemTagsCount)
}

func TestGetAllTag(t *testing.T) {
	setUpTestCase(t)
	// 注册路由
	tc := controller.TagController{}
	tc.RegisterRoutes(r.Group("/api"))
	// 创建一个用户
	user := &database.User{
		Email: "1@qq.com",
	}
	tx := database.DB.Create(user)
	if tx.Error != nil {
		t.Fatal("Create user failed:", tx.Error)
	}
	// 创建 5 条数据
	for i := 0; i < 5; i++ {
		tag := &database.Tag{
			Sign:   "⌚️",
			Name:   fmt.Sprintf("电子产品%d", i),
			Kind:   "expenses",
			UserID: user.ID,
		}
		database.DB.Create(tag)
	}
	// 初始化 w
	w := httptest.NewRecorder()
	// 发起请求
	req, _ := http.NewRequest(
		"GET",
		"/api/v1/tags",
		nil,
	)
	r.ServeHTTP(w, req)
	// 断言
	assert.Equal(t, 200, w.Code)
	// 解析响应
	var response api.GetAllTagResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("json.Unmarshal fail %v", err)
	}
	// 先用断言 后面会补充类型
	assert.Equal(t, 5, len(response.Resources))
}
