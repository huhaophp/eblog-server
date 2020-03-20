package routers

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/controllers/admin"
	"github.com/huhaophp/eblog/middleware/auth"
	"github.com/huhaophp/eblog/pkg/setting"
	"net/http"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	// 使用 Cookie 来存储
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("admin", store))
	// 设置请求 body 大小
	r.MaxMultipartMemory = 8 << 20
	// 设置静态资源地址
	r.Static("/assets", "./assets")
	// 设置模版路径
	r.LoadHTMLGlob("views/**/*")
	// 后台登陆
	r.Any("/admin/login", admin.Login)
	// 路由进行分组
	adminRoute := r.Group("admin")
	adminRoute.Use(auth.AdminSessionAuth())
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
	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "404 not found",
			"data": make(map[string]interface{}),
		})
	})

	return r
}