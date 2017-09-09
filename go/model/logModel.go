package model

// LogModel 日志数据模型
type LogModel struct {
	Type    string `json:"type"`
	Content string `json:"content"`
	Created int64  `json:"created"`
}

// TableName 表名
func (*LogModel) TableName() string {
	return "log"
}

// NewItems 实现 db.IModel 接口
func (*LogModel) NewItems() interface{} {
	items := new([]*LogModel)
	*items = make([]*LogModel, 0)
	return items
}
