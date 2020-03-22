package middleware

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AdminSessionAuth 后台登陆认证中间件
func AdminSessionAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if value := session.Get("admin"); value != nil {
			c.Next()
		} else {
			c.Redirect(http.StatusMovedPermanently, "/admin/login")
		}
	}
}
