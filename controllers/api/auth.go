package api

import (
	"github.com/gin-gonic/gin"
	r "github.com/huhaophp/eblog/controllers"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/pkg/util"
	"github.com/huhaophp/eblog/request/entity"
)

// Login 后台登陆
func Login(c *gin.Context) {
	var login entity.Login
	if parseErr := login.Parse(c); parseErr != nil {
		r.Json(c, 422, parseErr.Error(), gin.H{})
		return
	}
	if checkErr := login.Check(); checkErr != nil {
		r.Json(c, 422, checkErr.Error(), gin.H{})
		return
	}
	AuthModel := models.GetAuthByUsername(login.Username)
	if AuthModel.ID == 0 || AuthModel.Password != util.Md5(login.Password) {
		r.Json(c, 422, "账号或密码错误", gin.H{})
		return
	}
	ttl, token, err := util.GenerateToken(AuthModel.ID)
	if err != nil {
		r.Json(c, 422, "登陆错误", gin.H{})
	} else {
		r.Json(c, 0, "", gin.H{
			"ttl":   ttl,
			"token": token,
		})
	}
}
