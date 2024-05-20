package controller

import "github.com/gin-gonic/gin"

type ItemController struct{}

func (ctrl *ItemController) Get(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) Create(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) GetBanlance(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) Find(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) Destory(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
func (ctrl *ItemController) GetPaged(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) GetSummary(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("v1")
	// 注册路由
	v1.POST("/items", ctrl.Create)
	v1.GET("/items", ctrl.GetPaged)
	v1.GET("/items/balance", ctrl.GetBanlance)
	v1.GET("/items/summary", ctrl.GetSummary)
}
