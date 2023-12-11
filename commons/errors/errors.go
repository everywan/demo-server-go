package errors

import "fmt"

type ErrorCode struct {
	Code       int    `json:"code"`
	Msg        string `json:"msg"`
	InnerError error  `json:"inner_error"`
}

var _ error = (*ErrorCode)(nil)

func (e *ErrorCode) Error() string {
	msg := "error"
	if e.Code > 0 {
		msg += ", code: " + fmt.Sprint(e.Code)
	}
	if e.Msg != "" {
		msg += ", msg: " + e.Msg
	}
	if e.InnerError != nil {
		msg += ", inner_error: " + e.InnerError.Error()
	}
	return msg
}

func (e *ErrorCode) GetCode() int {
	return e.Code
}

func (e *ErrorCode) GetMsg() string {
	return e.Msg
}
