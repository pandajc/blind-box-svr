package routers

import "github.com/gin-gonic/gin"

func InitGinGroups(r *gin.Engine) {

	blindBox := r.Group("/blindBox")
	blindBox.GET("/orderList", getOrderList)

}
