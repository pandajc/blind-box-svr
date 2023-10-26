package global

var env string

func GetEnv() string {
	return env
}

func InitEnv(e string) {
	env = e
}