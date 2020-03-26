package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/huhaophp/eblog/pkg/setting"
	"github.com/unknwon/com"
	"time"
)

var jwtSecret = []byte(setting.AppSetting.JwtSecret)

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

// GenerateToken 生成令牌
func GenerateToken(id int) (string, int64, error) {
	jwtTokenTtl := com.StrTo(setting.AppSetting.JwtTokenTtl).MustInt64()
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + jwtTokenTtl,
			Issuer:    "eblog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecret); err != nil {
		return "", 0, err
	} else {
		return token, jwtTokenTtl, err
	}
}

//  ParseToken 解析令牌
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
