package model

import (
	"github.com/globalsign/mgo/bson"
	"time"
)

// Exam 测试题模型
type Exam struct {
	Id          bson.ObjectId `bson:"_id"`
	ExamName    string        `bson:"examName"`
	Description string        `bson:"description"`
	Deadline    time.Time     `bson:"deadline"`
	Size        int           `bson:"size"`
	Score       float64       `bson:"score"`

	/*
		Questions 测试题的题目
		内容格式：
		{
			"content": ["第一题", "第二题", "第三题"]
		}
	*/
	Questions string `bson:"questions"`

	/*
		Options 测试题的选项
		内容格式：
		{
			"content": [{
				"item": ["选项1的描述", "选项2的描述", "选项3的描述", "选项4的描述"]
			}, {
				"item": ["选项1的描述", "选项2的描述", "选项3的描述", "选项4的描述"]
			}, {
				"item": ["选项1的描述", "选项2的描述", "选项3的描述", "选项4的描述"]
			}]
		}
	*/
	Options string `bson:"options"`

	/*
		Keys 测试题的答案
		内容格式：
		{
			"content": ["A", "B", "C"]
		}
	*/
	Keys string `bson:"keys"`

	CreatedAt time.Time `bson:"createdAt"`
	UpdatedAt time.Time `bson:"updatedAt"`
}
