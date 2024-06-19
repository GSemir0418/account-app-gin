package controller

import (
	"account-app-gin/internal/api"
	"account-app-gin/internal/database"
	jwt_helper "account-app-gin/internal/jwt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SessionController struct{}

func (ctrl *SessionController) Create(c *gin.Context) {
	// 获取与校验请求体数据
	var body api.SessionRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.Error{Error: "Invalid request payload"})
		log.Print(err)
		return
	}
	// 查询验证码是否有效
	//SELECT * FROM validation_codes
	// WHERE
	// email=$1
	// AND
	// code=$2
	// AND
	// used_at IS NULL
	// ORDER BY created_at desc
	// LIMIT 1;
	var validationCode database.ValidationCode
	if result := database.DB.Where("email = ? AND code = ? AND used_at IS NULL", body.Email, body.Code).
		Order("created_at desc").
		First(&validationCode); result.Error != nil {
		c.JSON(http.StatusNotFound, api.Error{Error: "Code not found or already used"})
		log.Print(result.Error.Error())
		return
	}

	// 查询用户（无则创建）
	var user database.User
	if result := database.DB.FirstOrCreate(&user, database.User{Email: body.Email}); result.Error != nil {
		c.JSON(http.StatusNotFound, api.Error{Error: "Code not found or already used"})
		log.Print(result.Error.Error())
		return
	}
	// 返回 jwt
	jwt, err := jwt_helper.GenerateJWT(int(user.ID))
	if err != nil {
		log.Print("Generate JWT Error", err)
		c.JSON(http.StatusInternalServerError, api.Error{Error: "Failed to generate jwt"})
		return
	}

	res := api.SessionResponse{
		Jwt:    jwt,
		UserID: user.ID,
	}
	c.JSON(http.StatusOK, res)
}

func (ctrl *SessionController) RegisterRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.POST("/session", ctrl.Create)
}

func (ctrl *SessionController) Get(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *SessionController) Update(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *SessionController) GetPaged(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *SessionController) Find(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (ctrl *SessionController) Destory(c *gin.Context) {
	panic("not implemented") // TODO: Implement
}
