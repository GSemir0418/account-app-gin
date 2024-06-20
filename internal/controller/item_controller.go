package controller

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/database"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemController struct{}

func (ctrl *ItemController) Get(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ItemController) Create(c *gin.Context) {
	var body api.CreateItemRequest
	if err := c.BindJSON(&body); err != nil {
		log.Print(err)
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Invalid request payload"})
		return
	}

	// 检查 UserID 是否指向一个存在的 User 记录
	// var user database.User
	// if err := database.DB.First(&user, body.UserID).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// 查询 tag id 是否有效
	tagId := body.TagID
	var tag *database.Tag
	result := database.DB.First(&tag, tagId)
	if result.Error != nil {
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Can not find tag"})
		log.Print(result.Error.Error())
		return
	}
	// 中间件取出 user
	user, _ := c.Get("me")
	// 将 tags 放入 item 中
	var item database.Item
	item.UserID = user.(*database.User).ID
	item.TagID = tagId
	item.Amount = body.Amount
	item.HappenedAt = body.HappenedAt
	item.Kind = body.Kind

	if result := database.DB.Create(&item); result.Error != nil {
		log.Print(result.Error.Error())
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Failed to create item"})
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
	// 拿到 url 查询参数
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
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Get total count failed"})
		log.Print(err)
		return
	}

	// 查询分页的数据
	// 通过调用 Preload("Tag")，GORM 将在查询 Item 数据时自动执行额外的查询来加载每个 Item 的 Tag 数据，并将它们填充到返回的 Item 结构体的 Tag 字段。
	// 请注意 "Tag" 应该是你在 Item 结构体中定义的关联字段名称；如果你的字段名称不是 Tag，请用实际的字段名称替换它。
	if err := database.DB.Preload("Tag").Offset(offset).Limit(pageSize).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Get paged data failed"})
		log.Print(err)
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
