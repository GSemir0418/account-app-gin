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
		TagID:      tag.ID,
		UserID:     user.ID,
		Kind:       "in_come",
		HappenedAt: time.Now(),
	}
	tx = database.DB.Create(item)
	if tx.Error != nil {
		t.Fatal("Create item failed:", tx.Error)
	}
	// 测试此时两个表分别只有 1 条数据
	var tagCount int64
	database.DB.Model(&database.Tag{}).Count(&tagCount)
	var itemCount int64
	database.DB.Model(&database.Item{}).Count(&itemCount)
	assert.Equal(t, int64(1), tagCount)
	assert.Equal(t, int64(1), itemCount)
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
	// 断言删除成功
	assert.Equal(t, int64(0), tagCount)
	assert.Equal(t, int64(1), itemCount)
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

func TestGetTagSummaryWithMonth(t *testing.T) {
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
	// 创建 2 条 tag：电子产品1 和 电子产品2
	for i := 1; i < 3; i++ {
		tag := &database.Tag{
			Sign:   "⌚️",
			Name:   fmt.Sprintf("电子产品%d", i),
			Kind:   "expenses",
			UserID: user.ID,
		}
		database.DB.Create(tag)
	}
	// 分别给每个 tag 创建 2 条 item
	// 每个 tag 下的两条 item 时间不同
	item1 := &database.Item{
		Amount:     100,
		Kind:       "expenses",
		HappenedAt: time.Now().Add(time.Hour * 24 * 30),
		UserID:     user.ID,
		TagID:      1,
	}
	item2 := &database.Item{
		Amount:     100,
		Kind:       "expenses",
		HappenedAt: time.Now(),
		UserID:     user.ID,
		TagID:      1,
	}
	item3 := &database.Item{
		Amount:     200,
		Kind:       "expenses",
		HappenedAt: time.Now().Add(time.Hour * 24 * 30),
		UserID:     user.ID,
		TagID:      2,
	}
	item4 := &database.Item{
		Amount:     200,
		Kind:       "expenses",
		HappenedAt: time.Now(),
		UserID:     user.ID,
		TagID:      2,
	}
	items := []*database.Item{item1, item2, item3, item4}
	database.DB.Create(items)
	// 初始化 w
	w := httptest.NewRecorder()
	// 初始化当前月份
	month := time.Now().Format("2006-01")
	// 发起请求
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("/api/v1/tags/summary?month=%s", month),
		nil,
	)
	r.ServeHTTP(w, req)
	// 解析响应
	var response api.GetTagSummaryWithMonthResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("json.Unmarshal fail %v", err)
	}
	assert.Equal(t, 200, w.Code)
	// 断言
	// { resources: [ {id, name, sign, summary, kind}]}
	assert.Equal(t, 2, len(response.Resources))
	assert.Equal(t, "电子产品1", response.Resources[0].Name)
	assert.Equal(t, 100, response.Resources[0].Summary)
	assert.Equal(t, "电子产品2", response.Resources[1].Name)
	assert.Equal(t, 200, response.Resources[1].Summary)
}

func TestGetTagByID(t *testing.T) {
	setUpTestCase(t)
	// 注册路由
	tc := controller.TagController{}
	tc.RegisterRoutes(r.Group("/api"))
	// 创建一个用户
	user := &database.User{
		Email: "1@qq.com",
	}
	database.DB.Create(user)
	// 创建一个 tag
	tag := &database.Tag{
		Sign:   "⌚️",
		Name:   "电子产品",
		Kind:   "expenses",
		UserID: user.ID,
	}
	database.DB.Create(tag)
	// tag 下创建 5 条数据
	for i := 0; i < 5; i++ {
		item := &database.Item{
			Amount:     100,
			Kind:       "expenses",
			HappenedAt: time.Now(),
			UserID:     user.ID,
			TagID:      tag.ID,
		}
		database.DB.Create(item)
	}
	// 初始化 w
	w := httptest.NewRecorder()
	// 发起请求
	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("/api/v1/tags/%d", tag.ID),
		nil,
	)
	r.ServeHTTP(w, req)
	// 处理响应体
	var response database.Tag
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("json.Unmarshal fail %v", err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, tag.ID, response.ID)
	assert.Equal(t, tag.Sign, response.Sign)
	assert.Equal(t, tag.Name, response.Name)
	assert.Equal(t, 5, len(response.Items))
	assert.Equal(t, 100, response.Items[0].Amount)
}
