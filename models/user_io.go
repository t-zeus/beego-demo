package models

// 注册表单结构
type RegisterForm struct {
	Phone    string `form:"phone"`
	Name     string `form:"name"`
	Password string `form:"password"`
}
