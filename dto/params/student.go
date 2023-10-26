package params

type StudentReq struct {
	Name    string `form:"name" json:"name" binding:"required"`
	Age     int    `form:"age" json:"age" `
	PageNum int64  `form:"pageNum" json:"pageNum" `
}

type StudentCreateReq struct {
	Name  string `form:"name" json:"name" binding:"required"`
	Age   int    `form:"age" json:"age" `
	Phone string `form:"phone" json:"phone" `
}
