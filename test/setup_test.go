package test

import (
	"account-app-gin/config"
	"account-app-gin/internal/database"
	jwt_helper "account-app-gin/internal/jwt"
	"context"
	"net/http"
	"testing"

	"account-app-gin/internal/middleware"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
	c context.Context
)

func setUpTestCase(t *testing.T) {
	// 加载环境变量
	config.LoadEnv()
	// 连接数据库
	database.ConnectDB()

	r = gin.New()
	// 应用中间件
	r.Use(middleware.Me([]string{"/api/v1/session", "/api/v1/validation-codes", "/ping"}))
	// 默认上下文
	c = context.Background()
	// 清空 User 表 TRUNCATE 相当于先 DROP 在 CREATE
	// 存在外键约束的话，为了保证数据完整性，不能使用 TRUNCATE
	// 使用 DELETE 性能很差，所以这里使用更激进的策略，先关闭外键约束检查，清空后再恢复
	database.TruncateTables(t, []string{"users", "items", "tags", "validation_codes"})
}

func strPtr(s string) *string {
	return &s
}

func logIn(userID uint, req *http.Request) {
	jwtString, _ := jwt_helper.GenerateJWT(int(userID))
	req.Header = http.Header{
		"Authorization": []string{"Bearer " + jwtString},
	}
}
