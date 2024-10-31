package main

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

// 对称加密算法的JWT
func main() {
	m := map[string]any{
		"jti":      "jwt_1",             // 【保留字段】(JWT ID)是JWT的唯一标识
		"exp":      float64(3516239022), // 【保留字段】(Expiration time)过期时间(类型是float64、json.Number)
		"nbf":      float64(1516239022), // 【保留字段】(Not Before)这个参数指示了在什么时间之前，JWT是无效的
		"iat":      float64(1516239022), // 【保留字段】(Issued at)签发时间
		"aud":      "张三",                // 【保留字段】(Audience)接收对象
		"iss":      "jock.com",          // 【保留字段】(Issuser)签发主体
		"sub":      "lee",               // 【保留字段】(Subject)代表这个JWT的主体，即它的所有人
		"userName": "jock",              // 【业务】数据
		"userAge":  12,                  // 【业务】用户数据
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
	optionsFuncList := append(make([]jwt.ParserOption, 0),
		jwt.WithIssuedAt(),
		jwt.WithIssuer("jock.com"),
		jwt.WithAudience("张三"),
		jwt.WithSubject("lee"),
		jwt.WithValidMethods([]string{"HS256"}),
	)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}
		return hmacSecret, nil
	}, optionsFuncList...)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	// fmt.Println(token.Header)
	return token.Claims.(jwt.MapClaims), nil
}
