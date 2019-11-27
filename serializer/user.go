package serializer

import (
	"easy_learning/model"
	"time"
)

// User 用户序列化器
type User struct {
	Id        string    `json:"id"`
	UserName  string    `json:"username"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		Id:        user.Id.Hex(),
		UserName:  user.Username,
		Type:      user.Type,
		CreatedAt: user.CreatedAt,
	}
}
