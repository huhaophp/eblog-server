package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	// Upload the file to specific dst.
	err := c.SaveUploadedFile(file, "/static/uploadfile/"+file.Filename)
	if err != nil {
		c.String(http.StatusBadRequest, "保存失败 Error:%s", err.Error())
		return
	}
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
