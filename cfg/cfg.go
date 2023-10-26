package cfg

var Config = struct {
	Server struct {
		Port int
	}

	DB struct {
		Url          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Log struct {
		Level string
	}
}{}
