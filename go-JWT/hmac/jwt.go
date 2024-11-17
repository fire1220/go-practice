package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

const hmacKey = "helloKey"

type CustomClaimsUser struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

type customClaims struct {
	CustomClaims CustomClaimsUser
	jwt.RegisteredClaims
}

var _customClaims = new(customClaims)

func GetJWT() *customClaims {
	return _customClaims
}

func (c *customClaims) Encode(p CustomClaimsUser, base ...jwt.RegisteredClaims) (string, error) {
	c.CustomClaims = p
	if len(base) > 0 {
		_customClaims.RegisteredClaims = base[0]
	} else {
		_customClaims.RegisteredClaims = jwt.RegisteredClaims{
			Issuer:    "project",
			Subject:   "jock",
			Audience:  []string{"all"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(20 * time.Second)),
			NotBefore: nil,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.NewString(),
		}
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(hmacKey))
	if err != nil {
		return "", err
	}
	c.CustomClaims = CustomClaimsUser{}
	return tokenString, err
}

func (c *customClaims) Decode(tokenString string) (customClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(hmacKey), nil
	})
	if err != nil {
		return *c, err
	}
	if v, ok := token.Claims.(*customClaims); ok {
		return *v, nil
	} else {
		return *v, errors.New("type error")
	}
}
