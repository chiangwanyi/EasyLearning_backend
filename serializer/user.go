package serializer

import (
	"easy_learning/model"
	"github.com/globalsign/mgo/bson"
	"time"
)

// User 用户序列化器
type User struct {
	Id         string          `json:"id"`
	Username   string          `json:"username"`
	Email      string          `json:"email"`
	Type       string          `json:"type"`
	ClassList  []bson.ObjectId `json:"class_list"`
	InternalId string          `json:"internal_id"`
	Gender     string          `json:"gender"`
	SchoolName string          `json:"school_name"`
	CreatedAt  time.Time       `json:"created_at"`
}

// UserHome 用户首页序列化器
type UserHome struct {
	Id               string `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	Type             string `json:"type"`
	TeacherName      string `json:"teacher_name"`
	Classname        string `json:"classname"`
	ClassDescription string `json:"class_description"`
	Gender           string `json:"gender"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		Id:         user.Id.Hex(),
		Username:   user.Username,
		Email:      user.Email,
		Type:       user.Type,
		ClassList:  user.ClassList,
		InternalId: user.InternalId,
		Gender:     user.Gender,
		SchoolName: user.SchoolName,
		CreatedAt:  user.CreatedAt,
	}
}

func BuildStudents(items []model.User) (list []User) {
	for _, val := range items {
		list = append(list, BuildUser(val))
	}
	return list
}

func BuildUserHome(user model.User, class model.Class) UserHome {
	teacher, _ := model.FindUserById(class.TeacherId.Hex())
	return UserHome{
		Id:               user.Id.Hex(),
		Username:         user.Username,
		Email:            user.Email,
		Gender:           user.Gender,
		Type:             user.Type,
		TeacherName:      teacher.Username,
		Classname:        class.Classname,
		ClassDescription: class.Description,
	}
}
