package decorators

// 装饰结构体, 使结构体支持配置 Code,InnerError.
// 典型场景: service 接口需要返回失败状态码信息.
type CodeMsgInterface interface {
	GetCode() int
	GetMsg() string
	GetInnerError() error
}

type codemsg struct {
	code       int
	msg        string
	innerError error
}

func NewCodemsg(code int, msg string, innerError error) *codemsg {
	return &codemsg{
		code:       code,
		msg:        msg,
		innerError: innerError,
	}
}

var _ CodeMsgInterface = new(codemsg)

func (d *codemsg) GetCode() int {
	if d == nil {
		return 0
	}
	return d.code
}

func (d *codemsg) GetMsg() string {
	if d == nil {
		return ""
	}
	return d.msg
}

func (d *codemsg) GetInnerError() error {
	if d == nil {
		return nil
	}
	return d.innerError
}

func (d *codemsg) GetError() error {
	if d == nil {
		return nil
	}
	return d.innerError
}
