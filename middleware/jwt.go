package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/pkg/e"
	"github.com/huhaophp/eblog/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = e.SUCCESS
		authorization := c.GetHeader("Authorization")
		if authorization == "" {
			code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
		} else {
			bearers := strings.Split(authorization, " ")
			if bearers[1] == "" {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				claims, err := util.ParseToken(bearers[1])
				if err != nil {
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				} else if time.Now().Unix() > claims.ExpiresAt {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				}
			}
		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
