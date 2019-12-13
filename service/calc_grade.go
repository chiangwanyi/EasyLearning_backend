package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
	"encoding/json"
	"fmt"
	"github.com/globalsign/mgo/bson"
)

type CalcGrade struct {
	Answer map[string]interface{} `json:"answer" binding:"required"`
	ExamId string                 `json:"exam_id" binding:"required"`
}

func (service CalcGrade) CalcExamGrade(sid string, cid string) serializer.Response {
	if exam, err := model.FindExamById(service.ExamId); err == nil {
		key := make(map[string]interface{})
		_ = json.Unmarshal([]byte("{\"data\":"+exam.Keys+"}"), &key)

		count := 0
		for index, value := range key["data"].([]interface{}) {
			a := service.Answer["data"].([]interface{})[index].(map[string]interface{})["content"]
			k := value.(map[string]interface{})["content"]
			if a == k {
				fmt.Printf("第%d题正确，正确答案：%s\n", index+1, k)
				count++
			} else {
				fmt.Printf("第%d题错误，正确答案：%s，你的答案：%s\n", index+1, k, a)
			}
		}
		fmt.Println("总分：", float64(count)*exam.Score)

		gradebook := model.Gradebook{
			ClassId:    bson.ObjectIdHex(cid),
			ExamId:     bson.ObjectIdHex(service.ExamId),
			StudentId:  bson.ObjectIdHex(sid),
			FinalScore: float64(count) * exam.Score,
			Original:   service.Answer,
		}

		_ = gradebook.CreateGradeBook()

		return serializer.Response{
			Status: serializer.OK,
			Data:   serializer.BuildGradeBook(gradebook),
			Msg:    "提交成功，考试结束",
			Error:  "",
		}
	} else {
		return serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "内部错误",
		}
	}
}
