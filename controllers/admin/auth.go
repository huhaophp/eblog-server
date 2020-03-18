package admin

import (
	"log"
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

func CreateAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	if !ok {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			c.JSON(http.StatusOK, gin.H{
				"code": e.INVALID_PARAMS,
				"msg":  e.GetMsg(e.INVALID_PARAMS),
				"data": data,
			})
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
	token, ttl, err := util.GenerateToken(admin.ID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR_AUTH_TOKEN,
			"msg":  e.GetMsg(e.ERROR_AUTH_TOKEN),
			"data": data,
		})
	} else {
		data["token"] = token
		data["token_ttl"] = ttl
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": data,
		})
	}
}

func GetAuth(c *gin.Context) {
	c.GetHeader("token")
}
