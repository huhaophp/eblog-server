package admin

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/pkg/util"
	"github.com/huhaophp/eblog/request"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authTmplPath string = "auth/login.tmpl"
	authRediPath string = "/admin/home"
)

// Login 后台面板登陆
func Login(c *gin.Context) {
	if method := strings.ToUpper(c.Request.Method); method == "GET" {
		c.HTML(http.StatusOK, authTmplPath, nil)
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	valid := request.AuthLogRequestValid(username, password)
	if valid.HasErrors() {
		for _, validErr := range valid.Errors {
			c.HTML(http.StatusOK, authTmplPath, gin.H{
				"error": validErr.Message,
			})
			return
		}
	}
	admin := models.CheckAuth(username)
	if admin.ID == 0 || admin.Password != util.Md5(password) {
		c.HTML(http.StatusOK, authTmplPath, gin.H{
			"error": "账号或密码错误",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("admin", admin.ID)
	if saveErr := session.Save(); saveErr != nil {
		c.HTML(http.StatusOK, authTmplPath, gin.H{
			"error": "系统内部错误",
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, authRediPath)
	}
}
