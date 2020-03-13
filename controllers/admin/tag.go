package admin

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/pkg/e"
	"github.com/huhaophp/eblog/pkg/setting"
	"github.com/huhaophp/eblog/pkg/util"
	"github.com/huhaophp/eblog/request"
	"github.com/unknwon/com"
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

// AddTag 新增文章标签
func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	createdBy := c.PostForm("created_by")

	valid := request.TagAddRequestValid(name, createdBy, state)
	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
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

// EditTag 修改标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	modifiedBy := c.PostForm("modified_by")

	valid := validation.Validation{}

	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}

			affected := models.EditTag(id, data)
			if affected == 1 {
				code = e.SUCCESS
			} else {
				code = e.ERROR
			}
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, gin.H{
				"code": e.INVALID_PARAMS,
				"msg":  err,
				"data": make(map[string]string),
			})
		}
		return
	}

	if !models.ExistTagByID(id) {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg":  e.GetMsg(e.ERROR_NOT_EXIST_TAG),
			"data": make(map[string]string),
		})
		return
	}

	if models.DeleteTag(id) == 1 {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  "删除成功",
			"data": make(map[string]string),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "删除失败",
			"data": make(map[string]string),
		})
	}
}
