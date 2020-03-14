package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/huhaophp/eblog/pkg/setting"
	"log"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)

type Claims struct {
	Id int `json:"id"`
	jwt.StandardClaims
}

// GenerateToken 生成令牌
func GenerateToken(id int) (string, int64, error) {
	sec, err := setting.Cfg.GetSection("app")
	if err != nil {
		log.Fatal(2, "Fail to get section 'app': %v", err)
	}
	jwtTokenTtl := sec.Key("JWT_TOKEN_TTL").MustInt64(3600)
	claims := Claims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + jwtTokenTtl,
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, jwtTokenTtl, err
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
