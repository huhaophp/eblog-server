package util

import (
	"github.com/gin-gonic/gin"
	s "github.com/huhao/eblog/pkg/setting"
	"github.com/unknwon/com"
)

// GetPage 获取分页
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * s.PageSize
	}

	return result
}
