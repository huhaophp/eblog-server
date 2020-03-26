package api

import (
	"log"
	"mime/multipart"
	"os"

	r "github.com/huhaophp/eblog/controllers"

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
	uploadPath := setting.AppSetting.UploadDir
	CreateDir(uploadPath)
	filePath := uploadPath + "/" + file.Filename
	if saveErr := c.SaveUploadedFile(file, filePath); saveErr != nil {
		r.Json(c, 422, "上传失败", data)
	} else {
		data["file_path"] = filePath
		r.Json(c, 0, "上传成功", data)
	}
}

// CreateDir 创建目录
func CreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if mkdirAllErr := os.MkdirAll(path, os.ModePerm); mkdirAllErr != nil {
			log.Println(mkdirAllErr)
		}
		if chmodErr := os.Chmod(path, os.ModePerm); chmodErr != nil {
			log.Println(chmodErr)
		}
	}
}

// IsSupportedFileTypes 文件类型是否支持
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
