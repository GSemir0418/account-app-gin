package test

import (
	"account-app-gin/internal/database"
	"context"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
	c context.Context
)

func setUpTestCase(t *testing.T) {
	// 连接数据库
	database.ConnectDB()

	r = gin.New()
	// 默认上下文
	c = context.Background()
	// 清空 User 表
	if err := database.DB.Migrator().DropTable(&database.User{}); err != nil {
		t.Fatal(err)
	}
	// 清空 Items 表
	if err := database.DB.Migrator().DropTable(&database.User{}); err != nil {
		t.Fatal(err)
	}
}
