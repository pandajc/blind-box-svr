package routers

import (
	"fengshen-svr/biz"
	"fengshen-svr/common"
	"fengshen-svr/common/errs"
	"fengshen-svr/dto/params"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	req := &params.UserReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.Error(errs.NewRespErr(errs.ErrCodeInvalidParam, err.Error()))
		return
	}

	data, err := biz.GetUserBiz().Login(c, req)
	if err != nil {
		c.Error(err)
		return
	}
	common.SetGinRespData(c, data)
}
