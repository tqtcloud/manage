package errorx

const (
	defaultCode   = 1000 // 未定义错误
	parameterCode = 1001 // 请求参数错误
	userCode      = 1002 // 账号密码错误
)

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// NewCodeError 输入状态码生成错误
func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Msg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

// NewParameterError 请求参数错误
func NewParameterError(msg string) error {
	return NewCodeError(parameterCode, msg)
}

// NewUserError 账号密码错误
func NewUserError(msg string) error {
	return NewCodeError(userCode, msg)
}
