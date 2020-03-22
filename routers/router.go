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
	gin.SetMode(setting.RunMode)
	engine.MaxMultipartMemory = 8 << 20

	engine.POST("/admin/login", admin.Login)
	adminRoute := engine.Group("admin")
	adminRoute.Use(middleware.JWT())
	{
		// 后台首页
		adminRoute.GET("home", admin.Home)
		// 标签增删改查
		adminRoute.GET("tags", admin.TagIndex)
		adminRoute.GET("tags/create", admin.TagCreate)
		adminRoute.POST("tags/create", admin.TagCreate)
		adminRoute.GET("tags/edit", admin.TagEdit)
		adminRoute.POST("tags/edit", admin.TagEdit)
		adminRoute.POST("tags/delete", admin.TagDelete)
		// 文件上传
		adminRoute.POST("upload", admin.UploadFile)
	}
	// 404 页面
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "404 not found",
			"data": make(map[string]interface{}),
		})
	})
	return engine
}
