package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

// UserLoginService 用户登录服务
type UserLoginService struct {
	Email    string `json:"email" binding:"required,min=3,max=30"`
	Password string `json:"password" binding:"required,min=2,max=30"`
}

// Login 用户登录
func (service *UserLoginService) Login() (model.User, *serializer.Response) {
	// 通过 Email 查找用户
	if user, err := model.FindUserByEmail(service.Email); err == nil {
		// 检查密码
		if user.CheckPassword(service.Password) {
			return user, nil
		} else {
			return model.User{}, &serializer.Response{
				Data:  nil,
				Msg:   "",
				Error: "密码错误",
			}
		}
	} else {
		return model.User{}, &serializer.Response{
			Data:  err,
			Msg:   "",
			Error: "查询结果为空",
		}
	}
}
