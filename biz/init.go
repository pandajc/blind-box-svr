package biz

func InitAllBiz() error {
	NewStudentBiz()
	NewUserBiz()
	NewTelosScanBiz()
	return nil
}
