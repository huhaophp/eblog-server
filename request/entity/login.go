package entity

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 解析json参数
func (l *Login) Parse(c *gin.Context) error {
	if err := c.ShouldBindJSON(&l); err != nil {
		return err
	}
	log.Println(l)
	return nil
}

// 校验请求实体参数
func (l *Login) Check() error {
	valid := validation.Validation{}
	valid.Required(l.Username, "username").Message("用户名不能为空")
	valid.Required(l.Password, "password").Message("密码名不能为空")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return err
		}
	}
	return nil
}
