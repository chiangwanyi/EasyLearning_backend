package serializer

import (
	"easy_learning/model"
	"time"
)

// Class 班级序列化器
type Class struct {
	Id          string    `json:"id"`
	Classname   string    `json:"classname"`
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

// BuildClasses 序列化班级列表
func BuildClasses(items []model.Class) (classList []Class) {
	for _, val := range items {
		classList = append(classList, BuildClass(val))
	}
	return classList
}
