package errs

type RespErr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (r *RespErr) Error() string {
	return r.Msg
}

func NewRespErr(code int, msg string) *RespErr {
	return &RespErr{
		Code: code,
		Msg:  msg,
	}
}

const (
	Success             = 0
	ErrCodeSystemError  = 1000
	ErrCodeInvalidParam = 2000
)

var (
	ErrSystemError      = &RespErr{Code: ErrCodeSystemError, Msg: "system error"}
	ErrInvalidParamCode = &RespErr{Code: ErrCodeInvalidParam, Msg: "invalid param"}
)
