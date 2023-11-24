package codemsg

import "errors"

// error 类型定义
var (
	ErrorBadRequest = errors.New("bad request")
	ErrorNotFound   = errors.New("not found")
	ErrorForbidden  = errors.New("forbidden")
)
