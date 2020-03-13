package routers

import (
	"github.com/huhaophp/eblog/middleware/jwt"
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/controllers/admin"
	"github.com/huhaophp/eblog/pkg/setting"
)
// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.POST("/auth", admin.GetAuth)
	adminRoute := r.Group("/admin")
	adminRoute.Use(jwt.JWT())
	{
		adminRoute.GET("/tags", admin.GetTags)
		adminRoute.POST("/tags", admin.AddTag)
		adminRoute.PUT("/tags/:id", admin.EditTag)
		adminRoute.DELETE("/tags/:id", admin.DeleteTag)

		adminRoute.GET("articles", admin.GetArticles)
		adminRoute.GET("articles/:id", admin.GetArticle)
		adminRoute.POST("articles", admin.AddArticle)
		adminRoute.PUT("/articles/:id", admin.EditArticle)
		adminRoute.DELETE("/articles/:id", admin.DeleteArticle)
	}

	return r
}
