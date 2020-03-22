package admin

import (
	"github.com/gin-gonic/gin"
	r "github.com/huhaophp/eblog/controllers"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/pkg/util"
	"github.com/huhaophp/eblog/request"
)

// Login 后台登陆
func Login(c *gin.Context) {
	name := c.PostForm("username")
	pass := c.PostForm("password")
	data := gin.H{}
	valid := request.AuthLogRequestValid(name, pass)
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			r.Json(c, 422, err.Message, data)
		}
		return
	}
	AuthModel := models.GetAuthByUsername(name)
	if AuthModel.ID == 0 || AuthModel.Password != util.Md5(pass) {
		r.Json(c, 422, "账号或密码错误", data)
		return
	}
	ttl, token, err := util.GenerateToken(AuthModel.ID)
	if err != nil {
		r.Json(c, 422, "登陆错误", data)
	} else {
		data["ttl"] = token
		data["token"] = ttl
		r.Json(c, 0, "", data)
	}
}
