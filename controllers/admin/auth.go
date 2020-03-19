package admin

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/pkg/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	// 模版地址
	tmplPath string = "auth/login.tmpl"
	// 跳转地址
	rediPath string = "auth/login"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func AuthLogin(c *gin.Context) {
	if requestMethod := strings.ToUpper(c.Request.Method); requestMethod == "GET" {
		c.HTML(http.StatusOK, tmplPath, nil)
		return
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	if ok, _ := valid.Valid(&a); !ok {
		for _, validError := range valid.Errors {
			c.HTML(http.StatusOK, tmplPath, gin.H{
				"error": validError.Message,
			})
			return
		}
	}
	admin := models.CheckAuth(username)
	if admin.ID == 0 {
		c.HTML(http.StatusOK, tmplPath, gin.H{
			"error": "账号密码错误",
		})
		return
	}
	if admin.Password != util.Md5(password) {
		c.HTML(http.StatusOK, tmplPath, gin.H{
			"error": "账号密码错误",
		})
		return
	}
	session := sessions.Default(c)
	session.Set("admin", admin.ID)
	if saveErr := session.Save(); saveErr != nil {
		c.HTML(http.StatusOK, tmplPath, gin.H{
			"error": "系统服务错误",
		})
	} else {
		c.Redirect(http.StatusMovedPermanently, rediPath)
	}
}
