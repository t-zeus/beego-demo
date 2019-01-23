package routers

import (
	"beego-demo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	// 使用 beego 的 namespace 注册路由

	// 首先创建 /v1 命名空间
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/users",

			// 注册:POST
			beego.NSRouter("/register", &controllers.UserController{}, "post:Register"),

		),
	)

	// 注册 *Namespace
	beego.AddNamespace(ns)
}
