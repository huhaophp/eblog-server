package admin

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/pkg/e"
	"github.com/huhaophp/eblog/pkg/setting"
)

var fileTypes = []string{"image/jpeg", "image/png", "image/jpg"}

// UploadFile 上传文件
// @params file
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	if supported := IsSupportedFileTypes(file); !supported {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "不支持的文件类型",
			"data": make(map[string]string),
		})
		return
	}
	sec, err := setting.Cfg.GetSection("app")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "获取配置错误",
			"data": make(map[string]string),
		})
		return
	}

	dir := CreateDir(sec.Key("UPLOAD_DIR").String())
	if dir == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": e.ERROR,
			"msg":  "获取配置错误",
			"data": make(map[string]string),
		})
		return
	}
	if saveErr := c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", dir, file.Filename)); saveErr != nil {
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
func CreateDir(path string) string {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		if mkdirAllErr := os.MkdirAll(path, os.ModePerm); mkdirAllErr != nil {
			log.Println(mkdirAllErr)
		}
		if chmodErr := os.Chmod(path, os.ModePerm); chmodErr != nil {
			log.Println(chmodErr)
		}
	}
	return path
}

// isSupportedFileTypes 文件类型是否支持
func IsSupportedFileTypes(file *multipart.FileHeader) (supported bool) {
	supported = false
	fileType := file.Header.Get("Content-Type")
	for _, val := range fileTypes {
		if fileType == val {
			supported = true
			break
		}
	}
	return
}
