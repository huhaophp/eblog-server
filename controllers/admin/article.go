package admin

import (
	"github.com/gin-gonic/gin"
	r "github.com/huhaophp/eblog/controllers"
	"github.com/huhaophp/eblog/models"
	"github.com/huhaophp/eblog/request"
	"github.com/unknwon/com"
)

// ArticleIndex 标签列表
func ArticleIndex(c *gin.Context) {
	title := c.Query("title")
	state := com.StrTo(c.DefaultQuery("state", "-1")).MustInt()
	offset := com.StrTo(c.DefaultQuery("offset", "0")).MustInt()
	limit := com.StrTo(c.DefaultQuery("limit", "20")).MustInt()

	where := &models.Article{Title: title, State: state}
	total := models.GetArticlesTotal(where)
	articles := models.GetArticles(where, limit, offset)

	r.Json(c, 0, "", gin.H{
		"items": articles, "total": total,
	})
}

// ArticleAdd 文章添加
func ArticleAdd(c *gin.Context) {
	validErr, article := request.ArticleAddRequestValid(c)
	if validErr != nil {
		r.Json(c, 422, validErr.Error(), gin.H{})
		return
	}
	if addErr := models.AddArticle(&article); addErr != nil {
		r.Json(c, 422, addErr.Error(), gin.H{})
		return
	} else {
		r.Json(c, 0, "添加成功", gin.H{})
	}
}

// ArticleEdit 文章修改
func ArticleEdit(c *gin.Context) {
	validErr, article := request.ArticleAddRequestValid(c)
	if validErr != nil {
		r.Json(c, 422, validErr.Error(), gin.H{})
		return
	}
	id := com.StrTo(c.Param("id")).MustInt()
	if editErr := models.EditArticle(id, &article); editErr != nil {
		r.Json(c, 422, editErr.Error(), gin.H{})
	} else {
		r.Json(c, 0, "编辑文章成功", article)
	}
}

// ArticleDelete 文章删除
func ArticleDelete(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	if delErr := models.DelArticle(id); delErr != nil {
		r.Json(c, 422, delErr.Error(), gin.H{})
	} else {
		r.Json(c, 0, "删除成功", gin.H{})
	}
}
