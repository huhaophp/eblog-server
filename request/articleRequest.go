package request

import (
	"errors"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/models"
	"github.com/unknwon/com"
)

func ArticleAddRequestValid(c *gin.Context) (error, models.Article) {
	article := models.Article{}
	article.Title = c.PostForm("title")
	article.Cover = c.PostForm("cover")
	article.Desc = c.PostForm("desc")
	article.Content = c.PostForm("content")
	article.State = com.StrTo(c.PostForm("State")).MustInt()
	article.CateId = com.StrTo(c.PostForm("cate_id")).MustInt()

	valid := validation.Validation{}
	valid.Required(article.Title, "title").Message("文章名称不能为空")
	valid.Required(article.Cover, "cover").Message("文章封面不能为空")
	valid.Required(article.CateId, "cate_id").Message("文章所属栏目不能为空")
	valid.Required(article.Desc, "desc").Message("文章描述不能为空")
	valid.Required(article.Content, "state").Message("文章内容不能为空")
	valid.Range(article.State, 0, 1, "state").Message("文章状态只能是0和1")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return errors.New(err.Error()), article
		}
	}
	return nil, article
}
