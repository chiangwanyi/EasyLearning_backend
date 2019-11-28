package api

import (
	"easy_learning/config"
	"easy_learning/model"
	"easy_learning/serializer"
	"easy_learning/service"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHome 用户主页接口
func UserHome(c *gin.Context) {
	session := sessions.Default(c)

	fmt.Println(session.Get(config.SessionUserId))
	fmt.Println(session.Get(config.SessionClassId))

	if session.Get(config.SessionClassId) == nil {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.BadRequestError,
			Data:   nil,
			Msg:    "",
			Error:  "未设置当前课程",
		})
		return
	}

	uid := session.Get(config.SessionUserId).(string)
	cid := session.Get(config.SessionClassId).(string)

	if user, err := model.FindUserById(uid); err == nil {
		if class, err := model.FindClassById(cid); err == nil {
			c.JSON(http.StatusOK, serializer.Response{
				Status: serializer.OK,
				Data:   serializer.BuildUserHome(user, class),
				Msg:    "获取当前用户信息成功",
				Error:  "",
			})
		}
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "未找到用户，内部错误",
		})
	}
	return
}

// UserSetCurrentClass 用户设置当前班级接口
func UserSetCurrentClass(c *gin.Context) {

	s := service.UserSetCurrentClassService{}

	if err := c.ShouldBindJSON(&s); err == nil {
		classId := s.SetCurrentClass()

		session := sessions.Default(c)
		session.Set(config.SessionClassId, classId)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusOK, serializer.Response{
				Status: serializer.InternalServerError,
				Data:   err,
				Error:  "保存 Session 失败",
			})
			return
		}
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.OK,
			Data:   nil,
			Msg:    "设置课程成功",
			Error:  "",
		})
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "输入数据不匹配",
		})
	}
}

// UserLogout 用户登出接口
func UserLogout(c *gin.Context) {
	session := sessions.Default(c)

	// 删除 UserId
	session.Delete(config.SessionUserId)
	// 删除 ClassId
	session.Delete(config.SessionClassId)

	// 保存 Session
	if err := session.Save(); err != nil {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "保存 Session 失败",
		})
		return
	}
	c.JSON(http.StatusOK, serializer.Response{
		Status: serializer.OK,
		Data:   nil,
		Msg:    "登出成功",
		Error:  "",
	})
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	var s service.UserLoginService

	// 输入校验
	if err := c.ShouldBindJSON(&s); err == nil {
		if user, err := s.Login(); err == nil {

			// 保存 Session 状态
			session := sessions.Default(c)
			session.Clear()

			// 存储 UserId
			session.Set(config.SessionUserId, user.Id.Hex())
			// 初始化当前用户 ClassId
			session.Set(config.SessionClassId, "")

			if err := session.Save(); err != nil {
				c.JSON(http.StatusOK, serializer.Response{
					Status: serializer.InternalServerError,
					Data:   err,
					Error:  "保存 Session 失败",
				})
				return
			}
			c.JSON(http.StatusOK, serializer.Response{
				Status: serializer.OK,
				Data:   serializer.BuildUser(user),
				Msg:    "登录成功",
				Error:  "",
			})
		} else {
			c.JSON(http.StatusOK, err)
		}
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "输入数据不匹配",
		})
	}
}

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	var s service.UserRegisterService

	// 进行输入数据校验
	if err := c.ShouldBindJSON(&s); err == nil {
		// 注册
		if user, err := s.Register(); err != nil {
			c.JSON(http.StatusOK, err)
		} else {
			c.JSON(http.StatusOK, serializer.Response{
				Status: serializer.OK,
				Data:   serializer.BuildUser(user),
				Msg:    "注册成功",
				Error:  "",
			})
		}
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.BadRequestError,
			Data:   err,
			Msg:    "",
			Error:  "输入数据格式不匹配",
		})
	}
}
