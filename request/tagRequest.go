package request

import "github.com/astaxie/beego/validation"

func TagAddRequestValid(name string, state int) validation.Validation {
	valid := validation.Validation{}
	valid.Required(name, "name").Message("标签名称不能为空")
	valid.MaxSize(name, 100, "name").Message("标签名称最长为100字符")
	valid.Range(state, 0, 1, "state").Message("标签状态只允许0或1")
	return valid
}
