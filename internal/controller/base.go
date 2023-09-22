package controller

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
		// LastID // 可以按需添加 LastID
		Limit   int    `json:"limit"`
		Offset  int    `json:"offset"`
		Order   string `json:"order"`    // default desc.
		OrderBy string `json:"order_by"` // default id
	}
	// 作为 Response.Data
	PageBaseResponse struct {
		Total   int         `json:"total"`
		Records interface{} `json:"records"`
	}
)

// --------------------- request -------------------
type IDRequest struct {
	ID int64 `json:"id"`
}

type IDRequestUint struct {
	ID uint64 `json:"id"`
}

type IDRequestString struct {
	ID string `json:"id"`
}
