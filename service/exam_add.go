package service

import (
	"easy_learning/model"
	"easy_learning/serializer"
	"time"
)

type ExamAddServices struct {
	ExamName    string  `json:"exam_name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Deadline    string  `json:"deadline" binding:"required"`
	Size        int     `json:"size" binding:"required"`
	Score       float64 `json:"score" binding:"required"`
	Questions   string  `json:"questions" binding:"required"`
	Options     string  `json:"options" binding:"required"`
	Keys        string  `json:"keys" binding:"required"`
}

func (service *ExamAddServices) AddExam() (model.Exam, *serializer.Response) {
	deadline, _ := time.ParseInLocation("2006/1/2 15:4:05", service.Deadline, time.Local)
	exam := model.Exam{
		ExamName:    service.ExamName,
		Description: service.Description,
		Deadline:    deadline,
		Size:        service.Size,
		Score:       service.Score,
		Questions:   service.Questions,
		Options:     service.Options,
		Keys:        service.Keys,
	}

	if err := exam.CreateExam(); err != nil {
		return exam, &serializer.Response{
			Status: serializer.InternalServerError,
			Data:   err,
			Msg:    "",
			Error:  "注册失败",
		}
	}

	return exam, nil
}
