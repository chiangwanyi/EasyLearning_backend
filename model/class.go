package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Class 班级模型
type Class struct {
	Id          bson.ObjectId   `bson:"_id"`
	TeacherId   bson.ObjectId   `bson:"teacherId"`
	ClassName   string          `bson:"classname"`
	Description string          `bson:"description"`
	StudentList []bson.ObjectId `bson:"studentList"`
	ExamList    []bson.ObjectId `bson:"examList"`
	CreatedAt   time.Time       `bson:"createdAt"`
	UpdatedAt   time.Time       `bson:"updatedAt"`
}
