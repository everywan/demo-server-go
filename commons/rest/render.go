package rest

// Response 接口返回结果有两个状态: 接口状态(http.code)和业务状态(data.code).
// 基础服务对接口状态进行监控, 非 200 变多会报警. 业务对业务状态进行监控
type (
	Response struct {
		Code    int         `json:"code"`
		Message string      `json:"msg"`
		Data    interface{} `json:"data"`
	}
)

// 请求成功的响应
func SucResponse(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

// 请求成功的响应
func FailResponse(code int, msg string) *Response {
	return &Response{
		Code:    code,
		Message: msg,
	}
}
