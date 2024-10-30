package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 对称加密算法的JWT
func main() {
	// 创建JWT
	hmacSampleSecret := []byte("abcd")
	token1 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo":  "bar",
		"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"name": "jock",
	})
	tokenString1, err := token1.SignedString(hmacSampleSecret)
	fmt.Println(tokenString1, err)

	// 验证JWT
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYW1lIjoiam9jayIsIm5iZiI6MTQ0NDQ3ODQwMH0.nlfyvduuBOLVWlsAeYL6g1YCj7VnqJQFxOhknb1h-_0"
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if !token.Valid {
		fmt.Println("Invalid token")
		return
	}
	// fmt.Printf("%#v\n", token)
	fmt.Println(token.Claims.(jwt.MapClaims))
}
