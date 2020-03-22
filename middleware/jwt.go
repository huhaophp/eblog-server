package middleware

import (
	"github.com/gin-gonic/gin"
	r "github.com/huhaophp/eblog/controllers"
	"github.com/huhaophp/eblog/pkg/util"
	"strings"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			r.Json(c, 401, "Token is invalid", nil)
			c.Abort()
			return
		}
		bearers := strings.Split(authorization, " ")
		if bearers[1] == "" {
			r.Json(c, 401, "Token is invalid", nil)
			c.Abort()
			return
		}
		claims, err := util.ParseToken(bearers[1])
		if err != nil {
			r.Json(c, 401, err.Error(), nil)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			r.Json(c, 401, "Token is invalid", nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
