package admin

import (
	"fmt"
	r "github.com/huhaophp/eblog/controllers"
	"log"
	"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/huhaophp/eblog/pkg/setting"
)

var fileTypes = []string{"image/jpeg", "image/png", "image/jpg"}

// UploadFile 上传文件
func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	data := gin.H{}
	if file == nil {
		r.Json(c, 422, "文件格式错误", data)
		return
	}
	if supported := IsSupportedFileTypes(file); !supported {
		r.Json(c, 422, "不支持的文件类型", data)
		return
	}
	sec, err := setting.Cfg.GetSection("app")
	if err != nil {
		r.Json(c, 422, "获取配置错误", data)
		return
	}

	dir := CreateDir(sec.Key("UPLOAD_DIR").String())
	if dir == "" {
		r.Json(c, 422, "获取配置错误", data)
		return
	}
	if saveErr := c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", dir, file.Filename)); saveErr != nil {
		r.Json(c, 422, "上传失败", data)
	} else {
		r.Json(c, 0, "上传成功", data)
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
