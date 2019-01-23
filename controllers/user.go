package controllers

import "beego"

type UserController struct {
	beego.Controller
}

func (c *UserController) Register() {
	c.Data["json"] = "hello"
	c.ServeJSON()
}
