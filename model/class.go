package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Class 班级模型
type Class struct {
	Id          bson.ObjectId `bson:"_id"`
	TeacherId   bson.ObjectId `bson:"teacherId"`
	ClassName   string        `bson:"classname"`
	Description string        `bson:"description"`

	/*
		StudentsList 班级的学生 Id 列表
		内容格式：[]
	*/
	StudentsList []bson.ObjectId `bson:"studentsList"`
	CreatedAt    time.Time       `bson:"createdAt"`
	UpdatedAt    time.Time       `bson:"updatedAt"`
}
