package serializer

import (
	"easy_learning/model"
	"time"
)

// User 用户序列化器
type Class struct {
	Id          string    `json:"id"`
	Classname   string    `json:"username"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

// BuildClass 序列化班级
func BuildClass(class model.Class) Class {
	return Class{
		Id:          class.Id.Hex(),
		Classname:   class.Classname,
		Description: class.Description,
		CreatedAt:   class.CreatedAt,
	}
}
