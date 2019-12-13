package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
	"regexp"
)

// UserRegisterService 用户注册服务
type UserRegisterService struct {
	UserName        string `json:"username" binding:"required,min=2,max=20"`
	InternalId      string `json:"internal_id" binding:"required"`
	Email           string `json:"email" binding:"required,min=3,max=30"`
	Password        string `json:"password" binding:"required,min=2,max=30"`
	PasswordConfirm string `json:"password_confirm" binding:"required,min=2,max=30"`
	Gender          string `json:"gender" binding:"required"`
	Type            string `json:"type" binding:"required"`
}

// Valid 处理非法
func (service *UserRegisterService) Valid() *serializer.Response {

	// 检查邮箱格式是否正确
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(service.Email) {
		return &serializer.Response{
			Status: serializer.BadRequestError,
			Msg:    "",
			Error:  "邮箱格式不正确",
		}
	}

	if service.Password != service.PasswordConfirm {
		return &serializer.Response{
			Status: serializer.BadRequestError,
			Msg:    "",
			Error:  "两次输入的密码不匹配",
		}
	}

	// FindUserByEmail 查找邮箱是否重复
	if user, err := model.FindUserByEmail(service.Email); err == nil {
		return &serializer.Response{
			Status: serializer.BadRequestError,
			Data:   serializer.BuildUser(user),
			Msg:    "",
			Error:  "邮箱已存在",
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		Username:   service.UserName,
		InternalId: service.InternalId,
		Email:      service.Email,
		Gender:     service.Gender,
		Type:       service.Type,
	}

	// 判断输入数据是否合法
	if err := service.Valid(); err != nil {
		return user, err
	}

	// 生成加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "密码加密失败",
		}
	}

	// 创建用户
	if err := user.CreateUser(); err != nil {
		return user, &serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "注册失败",
		}
	}

	return user, nil
}
