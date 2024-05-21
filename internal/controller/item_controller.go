package controller

import (
	"account-app-gin/internal/database"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ItemController struct{}

func (ctrl *ItemController) Get(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) Create(c *gin.Context) {
	var item database.Item
	if err := c.BindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 检查 UserID 是否指向一个存在的 User 记录
	var user database.User
	if err := database.DB.First(&user, item.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&item); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
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
