package middlewares

import (
	"blind-box-svr/common"
	"blind-box-svr/common/errs"
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func HandleError(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("panic error= %v", err)
			common.SetGinRespErr(c, errs.ErrSystemError)
		}
	}()
	c.Next()
	errors := c.Errors
	for _, e := range c.Errors {
		fmt.Println(e)
	}
	if len(errors) > 0 {
		err := errors[0].Err
		if re, ok := err.(*errs.RespErr); ok {
			common.SetGinRespErr(c, re)
		} else {
			// todo print stack
			log.Errorf("unknown err=%v", err)
			common.SetGinRespErr(c, errs.ErrSystemError)
		}
		return
	}

}
