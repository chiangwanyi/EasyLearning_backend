package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
)

// UserRegisterService 用户注册服务
type UserRegisterService struct {
	UserName        string `json:"username" binding:"required,min=2,max=20"`
	Email           string `json:"email" binding:"required,min=3,max=30"`
	Password        string `json:"password" binding:"required,min=2,max=30"`
	PasswordConfirm string `json:"password_confirm" binding:"required,min=2,max=30"`
	Type            string `json:"type" binding:"required"`
}

// Valid 处理非法
func (service *UserRegisterService) Valid() *serializer.Response {
	if service.Password != service.PasswordConfirm {
		return &serializer.Response{
			Msg:   "",
			Error: "两次输入的密码不匹配",
		}
	}

	// FindUserByEmail 查找邮箱是否重复
	if user, err := model.FindUserByEmail(service.Email); err == nil {
		return &serializer.Response{
			Data:  serializer.BuildUser(user),
			Msg:   "",
			Error: "邮箱已存在",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		UserName: service.UserName,
		Email:    service.Email,
		Type:     service.Type,
	}

	// 判断输入数据是否合法
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 生成加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Data:  err,
			Msg:   "",
			Error: "密码加密失败",
		}
	}

	// 创建用户
	if err := user.CreateUser(); err != nil {
		return user, &serializer.Response{
			Data:  err,
			Msg:   "",
			Error: "注册失败",
		}
	}

	return user, nil
}
