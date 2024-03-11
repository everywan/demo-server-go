package rest

// Response 接口返回结果有两个状态: 接口状态(http.code)和业务状态(data.code).
// 基础服务对接口状态进行监控, 非 200 变多会报警. 业务对业务状态进行监控
type Response struct {
	Code    int         `json:"code"` // 0 表示成功
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

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

type (
	PageBaseRequest struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		// LastID // 按需添加所需类型的 LastID
		Order   string `json:"order"`    // default desc.
		OrderBy string `json:"order_by"` // default id
	}
	PageBaseResponse struct {
		Total   int         `json:"total"`
		Records interface{} `json:"records"`
		LastID  interface{} `json:"last_id"`
	}
)

func (req *PageBaseRequest) LoadDefault() {
	if req.Limit == 0 {
		req.Limit = 10
	}
	if req.Order == "" {
		req.Order = "desc"
	}
	if req.OrderBy == "" {
		req.OrderBy = "id"
	}
}
