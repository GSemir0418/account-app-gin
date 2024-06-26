package jwt_helper

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getHmacSecret() []byte {
	// 从环境变量中获取数据库连接字符串
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET is not set")
	}
	return []byte(secret)
}

func GenerateJWT(user_id uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	secret := getHmacSecret()

	return token.SignedString(secret)
}

func ParseJWT(jwtString string) (*jwt.Token, error) {
	secret := getHmacSecret()

	return jwt.Parse(jwtString, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
}
