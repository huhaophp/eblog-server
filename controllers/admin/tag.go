package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/huhao/eblog/models"
	"github.com/huhao/eblog/pkg/e"
	"github.com/huhao/eblog/pkg/setting"
	"github.com/huhao/eblog/pkg/util"
	"github.com/huhao/eblog/request"
	"github.com/unknwon/com"
	"net/http"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	data["items"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章标签
func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	createdBy := c.PostForm("created_by")

	valid := request.TagAddRequestValid(name, createdBy, state)
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		if ! models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改文章标签
func EditTag(c *gin.Context) {
}

// 删除文章标签
func DeleteTag(c *gin.Context) {
}
