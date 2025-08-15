package utils

import (
	"Blug/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userId,
		"exp":        time.Now().Add(time.Hour * time.Duration(config.Config.Jwt.LifeSpan)).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Config.Jwt.Secret))
}
func TokenValid(c *fiber.Ctx) error {
	tokenString := ExtractToken(c)
	println("tokenString: ", tokenString)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])

		}
		return []byte(config.Config.Jwt.Secret), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(c *fiber.Ctx) string {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	return authHeader[7:]

}
func ExtractTokenID(c *fiber.Ctx) (int, error) {
	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Config.Jwt.Secret), nil
	})
	if err != nil {
		return -1, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	// 如果jwt有效，将user_id转换为浮点数字符串，然后再转换为 uint32
	if ok && token.Valid {
		uid, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return -1, err
		}
		return int(uid), nil
	}

	return -1, nil
}
