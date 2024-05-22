package controller

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/database"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TagController struct{}

func (ctrl *TagController) Get(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TagController) Create(c *gin.Context) {
	var tag database.Tag
	if err := c.BindJSON(&tag); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Invalid request payload"})
		log.Print(err)
		return
	}

	// 检查 UserID 是否指向一个存在的 User 记录 后面有登录中间件就不用了
	var user database.User
	if err := database.DB.First(&user, tag.UserID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&tag); result.Error != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Could not create tag"})
		log.Print(result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (ctrl *TagController) Update(c *gin.Context) {
	var tag database.Tag
	// 获取 url 路径参数
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.First(&tag, id).Error; err != nil {
		log.Print(err)
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Tag not found"})
		return
	}

	tag.ID = uint(id)
	// 更新的数据可以传一个或多个
	// 为了保证灵活性，这里使用指针单独定义接口入参类型
	var body api.UpdateTagRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Invalid request payload"})
		log.Print(err)
		return
	}

	if body.Name != nil {
		tag.Name = *body.Name
	}
	if body.Sign != nil {
		tag.Sign = *body.Sign
	}
	if body.Kind != nil {
		tag.Kind = *body.Kind
	}

	// Save 是一个组合函数。 如果保存值不包含主键，它将执行 Create，否则它将执行 Update (包含所有字段)。
	// result := database.DB.Save(&tag)
	// 改用 Update
	result := database.DB.Model(&tag).Updates(tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Could not update tag"})
		log.Print(result.Error.Error())
		return
	}

	// 返回更新后的完整数据
	c.JSON(http.StatusOK, tag)
}

func (ctrl *TagController) GetBanlance(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TagController) Find(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TagController) Destory(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
func (ctrl *TagController) GetPaged(c *gin.Context) {
	// 拿到请求参数
	// page := c.Request.URL.Query().Get("page")
	// page := c.DefaultQuery("page", "1")
	// pageSize := c.Request.URL.Query().Get("page_size")
	// pageSize := c.DefaultQuery("page_size", "10")
	// page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	// var items []database.Item
	// var total int64

	// // 分页设置
	// offset := (page - 1) * pageSize

	// // 首先得到总数，用于分页信息
	// if err := database.DB.Model(&database.Item{}).Count(&total).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "get model total failed"})
	// 	return
	// }

	// // 查询分页的数据
	// if err := database.DB.Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }

	// // 响应，包括分页的数据和总数
	// c.JSON(http.StatusOK, api.GetPagedResponse{
	// 	Resources: items,
	// 	Pager: api.Pager{
	// 		Total:    total,
	// 		Page:     int64(page),
	// 		PageSize: int64(pageSize),
	// 	},
	// })
}

func (ctrl *TagController) GetSummary(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TagController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.POST("/tags", ctrl.Create)
	v1.PATCH("/tags/:id", ctrl.Update)
	v1.DELETE("/tags/:id", ctrl.Destory)
	v1.GET("/tags", ctrl.GetPaged)
	v1.GET("/tags/:id", ctrl.Get)
}
