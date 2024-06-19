package router

import (
	"account-app-gin/internal/controller"

	"github.com/gin-gonic/gin"
)

// 加载Controllers的方法，返回 controller 包的 Controller 接口的切片
func loadControllers() []controller.Controller {
	return []controller.Controller{
		&controller.ItemController{},
		&controller.TagController{},
		&controller.ValidationCodeController{},
		&controller.SessionController{},
	}
}

func New() *gin.Engine {
	// 创建路由
	r := gin.Default()
	// 注册路由
	rg := r.Group("/api")
	for _, ctrl := range loadControllers() {
		ctrl.RegisterRoutes(rg)
	}
	return r
}
