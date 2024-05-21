package controller

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/database"
	"errors"
	"net/http"
	"strconv"

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
	// 拿到请求参数
	// page := c.Request.URL.Query().Get("page")
	// page := c.DefaultQuery("page", "1")
	// pageSize := c.Request.URL.Query().Get("page_size")
	// pageSize := c.DefaultQuery("page_size", "10")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	var items []database.Item
	var total int64

	// 分页设置
	offset := (page - 1) * pageSize

	// 首先得到总数，用于分页信息
	if err := database.DB.Model(&database.Item{}).Count(&total).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "get model total failed"})
		return
	}

	// 查询分页的数据
	if err := database.DB.Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	// 响应，包括分页的数据和总数
	c.JSON(http.StatusOK, api.GetPagedResponse{
		Resources: items,
		Pager: api.Pager{
			Total:    total,
			Page:     int64(page),
			PageSize: int64(pageSize),
		},
	})
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
