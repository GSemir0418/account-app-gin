package controller

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/database"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type TagController struct{}

func (ctrl *TagController) Get(c *gin.Context) {
	var tag database.Tag
	// 获取 url 路径参数
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DB.Preload("Items").First(&tag, id).Error; err != nil {
		log.Print(err)
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (ctrl *TagController) Create(c *gin.Context) {
	var body api.CreateTagRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Invalid request payload"})
		log.Print(err)
		return
	}

	// 检查 UserID 是否指向一个存在的 User 记录 后面有登录中间件就不用了
	// var user database.User
	// if err := database.DB.First(&user, body.UserID).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
	// 		return
	// 	}

	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// 看下中间件有没有把 user 信息注入到上下文中
	// user, exists := c.Get("me")
	// log.Println(user, exists)
	// 中间件取出 user
	user, _ := c.Get("me")
	// 创建 tag
	// 将 tagIds 放入 item 中
	var tag database.Tag
	tag.UserID = user.(*database.User).ID
	tag.Sign = body.Sign
	tag.Name = body.Name
	tag.Kind = body.Kind

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
	// 删除 tag，首先要删除关联表中的记录，再删除 tag 本身
	// 所以使用事务来保证删除操作的原子性
	// 获取 url 路径参数
	id, _ := strconv.Atoi(c.Param("id"))

	// 开始事务
	tx := database.DB.Begin()
	if tx.Error != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Failed to start transaction"})
		log.Print(tx.Error.Error())
		return
	}

	// // 步骤1: 从多对多关联表中删除
	// if err := tx.Exec("DELETE FROM item_tags WHERE tag_id = ?", id).Error; err != nil {
	// 	tx.Rollback() // 回滚事务
	// 	c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Invalid request params"})
	// 	log.Print(err)
	// 	return
	// }

	// 步骤2: 删除标签本身
	if err := tx.Delete(&database.Tag{}, id).Error; err != nil {
		tx.Rollback() // 回滚事务
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Invalid request params"})
		log.Print(err)
		return
	}

	// 提交事务
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Failed to commit transaction"})
		log.Print(err)
		return
	}

	c.Status(http.StatusOK)
}

func (ctrl *TagController) GetPaged(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *TagController) GetSummary(c *gin.Context) {
	month := c.DefaultQuery("month", time.Now().Format("2006-01"))
	// 根据用户输入的月份，获得该月份的第一天和最后一天
	firstDayOfMonth, _ := time.Parse("2006-01", month)
	lastDayOfMonth := firstDayOfMonth.AddDate(0, 1, -1)

	var TagSummaries []api.TagSummary
	// 	SELECT
	//     tags.id AS tag_id,
	//     tags.name,
	//     tags.sign,
	//     tags.kind,
	//     COALESCE(SUM(items.amount), 0) AS summary
	// FROM
	//     tags
	// LEFT JOIN items ON tags.id = items.tag_id
	//     AND items.happened_at >= 'firstDayOfMonth'
	//     AND items.happened_at <= 'lastDayOfMonth'
	// GROUP BY
	//     tags.id, tags.name, tags.sign, tags.kind;
	if err := database.DB.Table("tags").
		// tags.id as id 选择tags表的id字段，并给这个字段的返回值起了一个别名 id。
		// tags.name，tags.sign，和tags.kind 直接选择了这些字段。
		// COALESCE(SUM(items.amount), 0) as summary 计算了所有相关items的amount字段之和，如果没有任何项与给定的tag相关联，则返回0。
		Select("tags.id as id, tags.name, tags.sign, tags.kind, COALESCE(SUM(items.amount), 0) as summary").
		// 在tags和items表之间进行左外连接。条件是tags.id字段必须与items.tag_id字段匹配，并且items表中的happened_at字段必须在firstDayOfMonth和lastDayOfMonth参数指定的范围内。
		Joins("LEFT JOIN items ON tags.id = items.tag_id AND items.happened_at >= ? AND items.happened_at <= ?", firstDayOfMonth, lastDayOfMonth).
		// 指定了用来分组的字段，这是为了计算每个不同标签的items.amount之和。
		Group("tags.id, tags.name, tags.sign, tags.kind").
		// 执行查询并将结果填充到TagSummaries变量中。
		Scan(&TagSummaries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Failed to get tag summary"})
		log.Print(err.Error())
		return
	}
	// 将结果编码为JSON格式并写入响应体
	response := api.GetTagSummaryWithMonthResponse{
		Resources: TagSummaries,
	}
	// Send the response
	c.JSON(http.StatusOK, response)
}

func (ctrl *TagController) GetAll(c *gin.Context) {
	// 获取全部 tags 数据
	var tags []database.Tag

	if err := database.DB.Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Failed to get all tags"})
		log.Print(err.Error())
		return
	}

	c.JSON(http.StatusOK, api.GetAllTagResponse{
		Resources: tags,
	})
}

func (ctrl *TagController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.POST("/tags", ctrl.Create)
	v1.PATCH("/tags/:id", ctrl.Update)
	v1.DELETE("/tags/:id", ctrl.Destory)
	v1.GET("/tags", ctrl.GetAll)
	v1.GET("/tags/:id", ctrl.Get)
	// GET /tags/summary?month=YYYY-MM
	v1.GET("/tags/summary", ctrl.GetSummary)
}
