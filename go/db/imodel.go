package db

type IModel interface {
	TableName() string
	NewItems() interface{}
}
