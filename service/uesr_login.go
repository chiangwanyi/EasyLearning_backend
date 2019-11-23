package service

// UserLoginService 用户登录服务
type UserLoginService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=20"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=30"`
}

