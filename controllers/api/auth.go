package api

import (
	"github.com/gin-gonic/gin"
	r "github.com/huhaophp/eblog/controllers"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/pkg/util"
	"github.com/huhaophp/eblog/request"
)

// Login 后台登陆
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	data := gin.H{}
	err := request.AuthLogRequestValid(username, password)
	if err != nil {
		r.Json(c, 422, err.Error(), data)
		return
	}
	AuthModel := models.GetAuthByUsername(username)
	if AuthModel.ID == 0 || AuthModel.Password != util.Md5(password) {
		r.Json(c, 422, "账号或密码错误", data)
		return
	}
	ttl, token, err := util.GenerateToken(AuthModel.ID)
	if err != nil {
		r.Json(c, 422, "登陆错误", data)
	} else {
		data["ttl"] = ttl
		data["token"] = token
		r.Json(c, 0, "", data)
	}
}
