package controllers

import (
	"beego"
	"errors"
	"github.com/astaxie/beego/validation"
)

const (
	ErrInputData = "数据输出错误"
)

// 父级
type BaseController struct {
	beego.Controller
}

// 在父级上封装表单验证函数
func (c *BaseController) VerifyForm(v interface{}) error {
	valid := validation.Validation{}
	// Valid 的内容：结构体或结构体指针
	ok, err := valid.Valid(v)
	if err != nil {
		return err
	}
	// 参数验证不通过
	if !ok {
		text := ""
		for _, err := range valid.Errors {
			text += err.Key + ":" + err.Message + ";"
		}
		return errors.New(text)
	}
	return nil
}
