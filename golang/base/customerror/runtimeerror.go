package customerror

// RuntimeError 预定义错误
var (
	RuntimeError = New("RuntimeError")
	Unauthorized = New("Unauthorized") //未授权
)

/*
runtimeError 自定义运行时错误
*/
type runtimeError struct {
	//错误信息字符串
	Message string
}

func (e *runtimeError) Error() string {
	return e.Message
}

func New(text string) error {
	return &runtimeError{text}
}
