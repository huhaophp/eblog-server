package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/controllers/admin"
	"github.com/huhaophp/eblog/middleware"
	"github.com/huhaophp/eblog/pkg/setting"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Cors())
	gin.SetMode(setting.RunMode)
	engine.MaxMultipartMemory = 8 << 20

	engine.POST("/admin/login", admin.Login)
	adminRoute := engine.Group("admin")
	adminRoute.Use(middleware.JWT())
	{
		// 标签列表
		adminRoute.GET("/tags", admin.TagIndex)
		// 标签标签
		adminRoute.POST("/tags", admin.TagAdd)
		// 编辑标签
		adminRoute.PUT("/tags/:id", admin.TagEdit)
		// 删除标签
		adminRoute.DELETE("/tags/:id", admin.TagDelete)
		// 栏目列表
		adminRoute.GET("/cates", admin.CateIndex)
		// 添加栏目
		adminRoute.POST("/cates", admin.CateAdd)
		// 编辑栏目
		adminRoute.PUT("/cates/:id", admin.CateEdit)
		// 删除栏目
		adminRoute.DELETE("/cates/:id", admin.CateDelete)
		// 文章列表
		adminRoute.GET("/articles", admin.ArticleIndex)
		// 添加文章
		adminRoute.POST("/articles", admin.ArticleAdd)
		// 编辑文章
		adminRoute.PUT("/articles/:id", admin.ArticleEdit)
		// 删除文章
		adminRoute.DELETE("/articles/:id", admin.ArticleDelete)
		// 文件上传
		adminRoute.POST("/upload", admin.UploadFile)
	}
	// 404 页面
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "404 not found",
			"data": gin.H{},
		})
	})
	return engine
}
