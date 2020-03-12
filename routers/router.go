package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/huhao/eblog/controllers/admin"
	"github.com/huhao/eblog/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	adminRoute := r.Group("/admin")
	{
		// admin login
		adminRoute.POST("/login", admin.Login)
		// admin logout
		adminRoute.DELETE("/logout", admin.Logout)

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
