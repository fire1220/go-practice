package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// 对称加密算法的JWT
func main() {
	m := map[string]any{
		"foo":  "bar",
		"name": "jock",
		"exp":  float64(time.Now().Add(1 * time.Hour).Unix()),
	}
	sig, err := HMACSign([]byte("abcd"), m)
	fmt.Println("对称加密的JWT")
	fmt.Println("生成JWT：")
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJuYW1lIjoiam9jayJ9.LbcLniyRlKNOMvWuIYi_GsjsJkKvPWZe8i4CdqnWGrc
	fmt.Println(sig, err)
	fmt.Println("验证JWT：")
	v, err := HMACVerify([]byte("abcd"), sig)
	fmt.Println(v, err)
	fmt.Println("过期时间：")
	fmt.Println(v.GetExpirationTime())
	fmt.Println(v.GetNotBefore())
}

// HMACSign 创建JWT
func HMACSign(hmacSecret []byte, m map[string]any) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(m)).SignedString(hmacSecret)
}

// HMACVerify 验证JWT
func HMACVerify(hmacSecret []byte, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	// fmt.Println(token.Header)
	return token.Claims.(jwt.MapClaims), nil
}
