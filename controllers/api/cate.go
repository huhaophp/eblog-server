package api

import (
	"github.com/gin-gonic/gin"
	r "github.com/huhaophp/eblog/controllers"
	_ "github.com/huhaophp/eblog/docs"
	"github.com/huhaophp/eblog/models"
	"github.com/unknwon/com"
)

func CateIndex(c *gin.Context) {
	name := c.Query("name")
	cates := models.GetCates(name)

	r.Json(c, 0, "", cates)
}

func CateAdd(c *gin.Context) {
	name := c.PostForm("name")
	state := com.StrTo(c.PostForm("state")).MustInt()
	data := gin.H{}
	if name == "" {
		r.Json(c, 422, "参数错误", data)
		return
	}
	cate := &models.Cate{
		Name: name,
	}
	if exist := models.GetCate(cate); exist.ID > 0 {
		r.Json(c, 422, "栏目已存在", data)
		return
	}
	cate.State = state
	if row := models.AddCate(cate); row > 0 {
		r.Json(c, 0, "添加成功", data)
	} else {
		r.Json(c, 0, "添加失败", data)
	}
}

// CateEdit 栏目修改
func CateEdit(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	state := com.StrTo(c.PostForm("state")).MustInt()
	name := c.PostForm("name")
	data := gin.H{}
	if id == 0 || name == "" {
		r.Json(c, 422, "参数错误", data)
		return
	}
	tag := &models.Cate{
		ID: id,
	}
	if exist := models.GetCate(tag); exist.ID == 0 {
		r.Json(c, 422, "栏目不存在", data)
	} else {
		models.EditCate(id, name, state)
		r.Json(c, 0, "编辑成功", data)
	}
}

func CateDelete(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	data := gin.H{}
	if id == 0 {
		r.Json(c, 422, "参数错误", data)
		return
	}
	cate := &models.Cate{
		ID: id,
	}
	if exist := models.GetCate(cate); exist.ID == 0 {
		r.Json(c, 422, "栏目不存在", data)
	} else {
		models.DelCate(id)
		r.Json(c, 0, "删除成功", data)
	}
}
