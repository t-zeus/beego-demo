package controllers

import (
	"beego"
	"beego-demo/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Register() {
	form := models.RegisterForm{}
	if err := c.ParseForm(&form); err != nil {
		beego.Debug("解析注册表单：", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
	beego.Debug("解析表单内容：", &form)

	if err := c.VerifyForm(&form); err != nil {
		beego.Debug("验证表单：", err)
		c.Data["json"] = models.NewErrorInfo(ErrInputData)
		c.ServeJSON()
		return
	}
}
