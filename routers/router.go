package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/controllers/api"
	_ "github.com/huhaophp/eblog/docs"
	"github.com/huhaophp/eblog/middleware"
	"github.com/huhaophp/eblog/pkg/setting"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(middleware.Cors())
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.StaticFS("/static/uploadfile", http.Dir(setting.AppSetting.UploadDir))
	gin.SetMode(setting.ServerSetting.RunMode)
	engine.MaxMultipartMemory = 8 << 20
	engine.POST("/api/login", api.Login)
	apiRoute := engine.Group("api")
	apiRoute.Use(middleware.JWT())
	{
		// 标签列表
		apiRoute.GET("/tags", api.TagIndex)
		// 标签标签
		apiRoute.POST("/tags", api.TagAdd)
		// 编辑标签
		apiRoute.PUT("/tags/:id", api.TagEdit)
		// 删除标签
		apiRoute.DELETE("/tags/:id", api.TagDelete)
		// 栏目列表
		apiRoute.GET("/cates", api.CateIndex)
		// 添加栏目
		apiRoute.POST("/cates", api.CateAdd)
		// 编辑栏目
		apiRoute.PUT("/cates/:id", api.CateEdit)
		// 删除栏目
		apiRoute.DELETE("/cates/:id", api.CateDelete)
		// 文章列表
		apiRoute.GET("/articles", api.ArticleIndex)
		// 添加文章
		apiRoute.POST("/articles", api.ArticleAdd)
		// 编辑文章
		apiRoute.PUT("/articles/:id", api.ArticleEdit)
		// 删除文章
		apiRoute.DELETE("/articles/:id", api.ArticleDelete)
		// 文件上传
		apiRoute.POST("/upload", api.UploadFile)
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
