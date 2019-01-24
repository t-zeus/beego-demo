package models

// 注册表单结构
type RegisterForm struct {
	Phone    string `form:"phone" valid:"Required"`
	Name     string `form:"name" valid:"Required"`
	Password string `form:"password" valid:"Required"`
}
