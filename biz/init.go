package biz

func InitAllBiz() error {
	NewStudentBiz()
	NewUserBiz()
	return nil
}
