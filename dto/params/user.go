package params

type UserReq struct {
	GameTime int64       `form:"gameTime" json:"gameTime" `
	Money    int64       `form:"money" json:"money"`
	BizData  interface{} `form:"bizData" json:"bizData"`
}

type UserResp struct {
	ID       int64       `form:"id" json:"id" `
	GameTime int64       `form:"gameTime" json:"gameTime" `
	Money    int64       `form:"money" json:"money"`
	Update   bool        `form:"update" json:"update"`
	BizData  interface{} `form:"bizData" json:"bizData"`
}
