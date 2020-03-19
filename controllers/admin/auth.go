package admin

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/huhaophp/eblog/pkg/logging"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/pkg/e"
	"github.com/huhaophp/eblog/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func AuthLogin(c *gin.Context) {
	requestMethod := c.Request.Method
	if requestMethod == "GET" {
		c.HTML(http.StatusOK, "auth/login.tmpl", gin.H{
			"title": "Main website",
		})
	} else {
		username := c.PostForm("username")
		password := c.PostForm("password")
		valid := validation.Validation{}
		a := auth{Username: username, Password: password}
		ok, _ := valid.Valid(&a)
		data := make(map[string]interface{})
		if !ok {
			for _, err := range valid.Errors {
				logging.Info(err.Key, err.Message)
				return
			}
		}
		admin := models.CheckAuth(username)
		if admin.ID == 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": e.ERROR_AUTH,
				"msg":  e.GetMsg(e.ERROR_AUTH),
				"data": data,
			})
			return
		}
		if admin.Password != util.Md5(password) {
			c.JSON(http.StatusOK, gin.H{
				"code": e.ERROR_AUTH,
				"msg":  e.GetMsg(e.ERROR_AUTH),
				"data": data,
			})
			return
		}
		
		session := sessions.Default(c)
		session.Set("admin", admin.ID)
		if err := session.Save(); err != nil {
			logging.Info("session save error")
		} else {
			c.Redirect(http.StatusMovedPermanently, "/admin/home")
		}
	}
}
