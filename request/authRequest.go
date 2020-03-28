package request

import (
	"github.com/astaxie/beego/validation"
)

func AuthLogRequestValid(username, password string) error {
	valid := validation.Validation{}
	valid.Required(username, "username").Message("用户名不能为空")
	valid.MaxSize(username, 50, "username").Message("用户名最长为50字符")
	valid.Required(password, "password").Message("密码名不能为空")
	valid.MaxSize(password, 50, "password").Message("密码最长为50字符")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			return err
		}
	}
	return nil
}
