package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	signingKey = []byte("example_key")
)

func main() {
	// 创建JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = 123
	claims["username"] = "john_doe"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 设置过期时间为24小时

	// 签名JWT
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated JWT:", tokenString)

	// 验证JWT
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		fmt.Println("Error parsing token:", err)
		return
	}

	if !parsedToken.Valid {
		fmt.Println("Invalid token")
		return
	}

	claims = parsedToken.Claims.(jwt.MapClaims)
	fmt.Println("user_id:", claims["user_id"], "username:", claims["username"])
}
