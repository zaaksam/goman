package model

// ListModel 列表数据模型
type ListModel struct {
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	PageCount int         `json:"pageCount"`
	Total     int64       `json:"total"`
	Items     interface{} `json:"items"`
}
