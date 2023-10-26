package routers

import (
	"fengshen-svr/biz"
	"fengshen-svr/common"
	"fengshen-svr/common/errs"
	"fengshen-svr/dto/params"

	"github.com/gin-gonic/gin"
)

func getStudentList(c *gin.Context) {
	req := &params.StudentReq{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Error(errs.NewRespErr(errs.ErrCodeInvalidParam, err.Error()))
		return
	}
	list, err := biz.GetStudentBiz().GetStudentList(req)
	if err != nil {
		c.Error(err)
		return
	}
	common.SetGinRespData(c, list)
}

func createStudent(c *gin.Context) {
	req := &params.StudentCreateReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(errs.NewRespErr(errs.ErrCodeInvalidParam, err.Error()))
		return
	}
	err = biz.GetStudentBiz().CreateStudent(req)
	if err != nil {
		c.Error(err)
		return
	}
	common.SetGinRespData(c, nil)
}