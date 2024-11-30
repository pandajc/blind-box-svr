package routers

import "github.com/gin-gonic/gin"

func InitGinGroups(r *gin.Engine) {
	st := r.Group("/student")
	st.GET("/list", getStudentList)
	st.POST("/create", createStudent)
	user := r.Group("/user")
	user.POST("/login", login)

	blindBox := r.Group("/blindBox")
	blindBox.GET("/orderList", getOrderList)

}
