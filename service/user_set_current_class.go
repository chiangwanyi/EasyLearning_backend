package service

type UserSetCurrentClassService struct {
	ClassId string `json:"class_id" binding:"required"`
}

func (service *UserSetCurrentClassService) SetCurrentClass() string  {
	return service.ClassId
}
