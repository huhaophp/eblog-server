package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/request"
	"github.com/unknwon/com"
	"net/http"
	"strings"
)

const (
	tagIndexTmplPath  string = "tags/index.tmpl"
	tagCreateTmplPath string = "tags/create.tmpl"
	tagCreateRediPath string = "/admin/tags"
	tagEditTmplPath   string = "tags/edit.tmpl"
)

// TagIndex 标签管理
func TagIndex(c *gin.Context) {
	name := c.Query("name")
	c.HTML(http.StatusOK, tagIndexTmplPath, gin.H{
		"tags": models.GetTags(name),
		"name": name,
	})
}

// TagCreate 标签创建
func TagCreate(c *gin.Context) {
	if method := strings.ToUpper(c.Request.Method); method == "GET" {
		c.HTML(http.StatusOK, tagCreateTmplPath, nil)
		return
	} else {
		name := c.PostForm("name")
		state := com.StrTo(c.PostForm("state")).MustInt()
		valid := request.TagAddRequestValid(name, state)
		if valid.HasErrors() {
			for _, validErr := range valid.Errors {
				c.HTML(http.StatusOK, tagCreateTmplPath, gin.H{
					"error": validErr.Message,
				})
				return
			}
		}
		if exist := models.ExistTagByName(name); exist {
			c.HTML(http.StatusOK, tagCreateTmplPath, gin.H{
				"error": "标签名已存在",
			})
			return
		}
		if success := models.AddTag(name, state, "admin"); success {
			c.Redirect(http.StatusMovedPermanently, tagCreateRediPath)
		} else {
			c.HTML(http.StatusOK, tagCreateTmplPath, gin.H{
				"error": "标签创建失败",
			})
		}
	}
}

// TagEdit 标签编辑
func TagEdit(c *gin.Context) {
	id := com.StrTo(c.Query("id")).MustInt()
	if method := strings.ToUpper(c.Request.Method); method == "GET" {
		tag := models.GetTag(id)
		c.HTML(http.StatusOK, tagEditTmplPath, gin.H{
			"tag": tag,
		})
		return
	}
	name := c.PostForm("name")
	state := com.StrTo(c.PostForm("state")).MustInt()
	valid := request.TagAddRequestValid(name, state)
	if valid.HasErrors() {
		for _, validErr := range valid.Errors {
			c.HTML(http.StatusOK, tagCreateTmplPath, gin.H{
				"error": validErr.Message,
			})
			return
		}
	}
	data := make(map[string]interface{})
	data["name"] = name
	data["state"] = state
	if success := models.EditTag(id, data); success {
		c.Redirect(http.StatusMovedPermanently, tagCreateRediPath)
	} else {
		c.HTML(http.StatusOK, tagEditTmplPath, gin.H{
			"error": "标签创建失败",
		})
	}
}

// TagDelete 标签删除
func TagDelete(c *gin.Context) {
	id := com.StrTo(c.PostForm("id")).MustInt()
	if id <= 0 {
		c.Redirect(http.StatusMovedPermanently, tagCreateRediPath)
	}
	if success := models.DeleteTag(id); success {
		c.Redirect(http.StatusMovedPermanently, tagCreateRediPath)
	} else {
		c.HTML(http.StatusOK, tagEditTmplPath, gin.H{
			"error": "标签创建失败",
		})
	}
}
