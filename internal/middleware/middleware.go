package middleware

import (
	"account-app-gin/internal/database"
	jwt_helper "account-app-gin/internal/jwt"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Me(whiteList []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		// 检测白名单
		for _, s := range whiteList {
			if has := strings.HasPrefix(path, s); has {
				c.Next()
				return
			}
		}
		user, err := getMe(c)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"error": err.Error(),
			})
			return
		}
		// 将 me 放到上下文中，作为全局变量
		c.Set("me", user)
		c.Next()
	}
}

func getMe(c *gin.Context) (database.User, error) {
	var user database.User

	auth := c.GetHeader("Authorization")
	if len(auth) < 8 {
		return user, fmt.Errorf("JWT is required")
	}
	// 截取 Bearer 后的字符
	jwtString := auth[7:]
	t, err := jwt_helper.ParseJWT(jwtString)
	if err != nil {
		return user, fmt.Errorf("invalid jwt")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return user, fmt.Errorf("invalid jwt")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return user, fmt.Errorf("invalid jwt")
	}

	if tx := database.DB.Find(&user, userID); tx.Error != nil {
		return user, fmt.Errorf("无效的jwt")
	}

	return user, nil
}
