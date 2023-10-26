package biz

import (
	"fengshen-svr/dto/params"
	"fengshen-svr/global"
	"fengshen-svr/model"

	log "github.com/sirupsen/logrus"
)

type IStudentBiz interface {
	GetStudentList(param *params.StudentReq) ([]*model.TStudent, error)
	CreateStudent(param *params.StudentCreateReq) error
}

type StudentBiz struct {
}

var studentBiz IStudentBiz

func NewStudentBiz() {
	studentBiz = &StudentBiz{}
}

func GetStudentBiz() IStudentBiz {
	return studentBiz
}

func (s *StudentBiz) GetStudentList(param *params.StudentReq) ([]*model.TStudent, error) {
	pageParam := model.NewPage(10, param.PageNum, model.BuildDescs("id")...)
	page, err := model.TStudentMgr(global.GetDB()).SelectPage(pageParam)
	if err != nil {
		log.Errorf("GetStudentList call SelectPage err=%v", err)
		return nil, err
	}
	records := page.GetRecords()
	students := records.([]model.TStudent)
	return ConvertStructToPointerInArr(students), nil
}

func (s *StudentBiz) CreateStudent(param *params.StudentCreateReq) error {
	tx := model.TStudentMgr(global.GetDB()).Create(param)
	err := tx.Error
	if err != nil {
		log.Errorf("CreateStudent call create err=%v", err)
		return err
	}
	return nil
}

func ConvertStructToPointerInArr[T interface{}](arr []T) []*T {
	var res []*T
	for _, v := range arr {
		res = append(res, &v)
	}
	return res
}
