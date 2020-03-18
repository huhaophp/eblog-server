package admin

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/pkg/e"
)

const uploadFilePath string = "./static/uploadfile"

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	CreateDir(uploadFilePath)
	saveUploadedFileErr := c.SaveUploadedFile(file, uploadFilePath+"/"+file.Filename)
	if saveUploadedFileErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "上传失败",
			"data": make(map[string]string),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": e.SUCCESS,
			"msg":  e.GetMsg(e.SUCCESS),
			"data": make(map[string]string),
		})
	}
}

// CreateDir 创建目录
func CreateDir(Path string) string {
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		os.MkdirAll(Path, os.ModePerm)
		os.Chmod(Path, os.ModePerm)
	}
	return Path
}
