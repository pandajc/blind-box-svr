package routers

import (
	"blind-box-svr/biz"
	"blind-box-svr/common"
	"blind-box-svr/common/errs"
	"blind-box-svr/dto/params"

	"github.com/gin-gonic/gin"
)

func getOrderList(c *gin.Context) {
	req := &params.OrderParam{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		c.Error(errs.NewRespErr(errs.ErrCodeInvalidParam, err.Error()))
		return
	}

	data, err := biz.GetTelosScanBiz().GetOrderList(req)
	if err != nil {
		c.Error(err)
		return
	}
	common.SetGinRespData(c, data)
}
