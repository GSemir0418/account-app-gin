package controller

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/database"
	"account-app-gin/internal/email"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ValidationCodeController struct{}

func (ctrl *ValidationCodeController) Create(c *gin.Context) {
	// 从请求体中解析 email
	var body api.CreateValidationCodeRequest
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Invalid request payload"})
		log.Print(err)
		return
	}
	// 生成随机四位验证码
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个1000到9999之间的随机数
	code := 1000 + r.Intn(9000)

	// 保存验证码到数据库
	var vc database.ValidationCode
	vc.Email = body.Email
	vc.Code = fmt.Sprint(code)
	if result := database.DB.Create(&vc); result.Error != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Could not create validation code"})
		log.Print(result.Error.Error())
		return
	}

	// 发送邮件
	err := email.SendValidationCode(vc.Email, vc.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Could not send validation code"})
		log.Print(err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (ctrl *ValidationCodeController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.POST("/validation-codes", ctrl.Create)
}

func (ctrl *ValidationCodeController) Get(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ValidationCodeController) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ValidationCodeController) GetPaged(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ValidationCodeController) Find(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *ValidationCodeController) Destory(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
