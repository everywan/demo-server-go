package rest

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
