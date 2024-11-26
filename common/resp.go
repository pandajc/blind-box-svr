package common

import (
	"blind-box-svr/common/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SetGinRespData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &Resp{
		Code: errs.Success,
		Msg:  "success",
		Data: data,
	})
}

func SetGinRespErr(c *gin.Context, err *errs.RespErr) {
	c.JSON(http.StatusOK, &Resp{
		Code: err.Code,
		Msg:  err.Msg,
	})
}
