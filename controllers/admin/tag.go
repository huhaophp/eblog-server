package admin

import (
	"github.com/gin-gonic/gin"
	r "github.com/huhaophp/eblog/controllers"
	"github.com/huhaophp/eblog/models"
	"github.com/unknwon/com"
)

// TagIndex 标签列表
func TagIndex(c *gin.Context) {
	name := c.Query("name")
	tags := models.GetTags(name)

	r.Json(c, 0, "", tags)
}

// TagAdd 新增标签
func TagAdd(c *gin.Context) {
	name := c.PostForm("name")
	data := gin.H{}
	if name == "" {
		r.Json(c, 422, "参数错误", data)
		return
	}
	tag := &models.Tag{
		Name: name,
	}
	if exist := models.GetTag(tag); exist.ID > 0 {
		r.Json(c, 422, "标签已存在", data)
		return
	}
	if row := models.AddTag(tag); row > 0 {
		r.Json(c, 0, "添加成功", data)
	} else {
		r.Json(c, 0, "添加成功", data)
	}
}

// TagAdd 新增标签
func TagEdit(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	data := gin.H{}
	if id == 0 || name == "" {
		r.Json(c, 422, "参数错误", data)
		return
	}
	tag := &models.Tag{
		ID: id,
	}
	if exist := models.GetTag(tag); exist.ID == 0 {
		r.Json(c, 422, "标签不存在", data)
		return
	}

	models.EditTag(id, name)

	r.Json(c, 0, "编辑成功", data)
}

// TagDelete 新增标签
func TagDelete(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	data := gin.H{}
	if id == 0 {
		r.Json(c, 422, "参数错误", data)
		return
	}
	tag := &models.Tag{
		ID: id,
	}
	if exist := models.GetTag(tag); exist.ID == 0 {
		r.Json(c, 422, "标签不存在", data)
		return
	}

	models.DelTag(id)

	r.Json(c, 0, "删除成功", data)
}

