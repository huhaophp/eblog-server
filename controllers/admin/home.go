package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Home 后台首页
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/admin.tmpl", gin.H{})
}
