package api

import (
	"easy_learning/config"
	"easy_learning/serializer"
	"easy_learning/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CurrentUser 返回当前登录的用户接口
func CurrentUser(c *gin.Context) {

}

// JoinClass 加入班级接口
func UserJoinClass(c *gin.Context) {
	var s service.UserJoinClassService

	if err := c.ShouldBindJSON(&s); err == nil {
		session := sessions.Default(c)
		if err := s.JoinClass(session.Get(config.SessionUserId).(string)); err != nil {
			c.JSON(http.StatusOK, err)
			return
		}
		c.JSON(http.StatusOK, serializer.Response{
			Status: serializer.OK,
			Data:   nil,
			Msg:    "加入班级成功",
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
