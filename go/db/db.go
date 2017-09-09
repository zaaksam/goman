package db

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/zaaksam/goman/go/model"
)

type IModel interface {
	TableName() string
	NewItems() interface{}
}

// Engine 数据库引擎
var Engine *xorm.Engine
var tables []xorm.TableName

func init() {
	tables = make([]xorm.TableName, 20)

	dir := path.Dir(os.Args[0])
	dir = filepath.ToSlash(dir) // .../goman/go
	file := filepath.Join(dir, "goman.db")

	_, err := os.Stat(file)
	if err != nil && os.IsNotExist(err) {
		_, err = os.Create(file)
		if err != nil {
			panic(fmt.Errorf("创建DB文件错误：%s", err))
		}
	}

	Engine, err = xorm.NewEngine("sqlite3", file)

	err = Engine.Ping()
	if err != nil {
		panic(fmt.Sprintf("DB连接不通：%s", err))
	}

	//结构体命名与数据库一致
	Engine.SetMapper(core.NewCacheMapper(new(core.SameMapper)))

	err = createTable(
		&model.LogModel{},
	)
	if err != nil {
		os.Remove(file)
		panic(err)
	}
}

func createTable(beans ...xorm.TableName) (err error) {
	// tables, err := Engine.DBMetas()
	// if err != nil {
	// 	panic(fmt.Sprintf("获取DB信息错误：%s", err))
	// }

	for i, l := 0, len(beans); i < l; i++ {
		err = Engine.CreateTables(beans[i])
		if err != nil {
			return fmt.Errorf("创建数据库表[%s]失败：%s", beans[i].TableName(), err)
		}
	}
	return
}

// NewSession 创建新的数据库操作对象
func NewSession() *xorm.Session {
	session := Engine.NewSession()
	session.IsAutoClose = true
	return session
}

// GetList 获取分页数据
func GetList(session *xorm.Session, md IModel, pageIndex, pageSize int) (list *model.ListModel, err error) {
	list = &model.ListModel{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Items:     md.NewItems(),
	}

	//克隆查询对象
	sessionCount := session.Clone()
	defer sessionCount.Close()

	//清空Orderby条件，在某些数据库count不能带orderby
	sessionCount.Statement.OrderStr = ""

	if groupBy := sessionCount.Statement.GroupByStr; groupBy != "" {
		sessionCount.Statement.GroupByStr = "" // count不能带groupby
		list.Total, err = sessionCount.Select("count(DISTINCT " + groupBy + ")").Count(md)
	} else {
		sessionCount.Select("") // select置空
		list.Total, err = sessionCount.Count(md)
	}
	if err != nil {
		return
	}

	//计算分页
	resetPagination(list)

	//没有数据，提前返回
	if list.Total <= 0 {
		return
	}

	recordIndex := (list.PageIndex - 1) * list.PageSize
	session.Limit(list.PageSize, recordIndex)

	err = session.Table(md).Find(list.Items)
	return
}

func resetPagination(list *model.ListModel) {
	//检查页面长度
	if list.PageSize <= 0 {
		list.PageSize = 10
	} else if list.PageSize > 500 {
		list.PageSize = 500
	}

	//计算总页数
	cnt := int(list.Total / int64(list.PageSize))
	mod := int(list.Total % int64(list.PageSize))
	if mod > 0 {
		cnt++
	}

	//检查页面索引
	switch {
	case cnt == 0:
		list.PageIndex = 1
	case list.PageIndex > cnt:
		list.PageIndex = cnt
	case list.PageIndex <= 0:
		list.PageIndex = 1
	}

	//设置页面总页数
	list.PageCount = cnt
}
