package db

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/zaaksam/goman/go/model"
	"xorm.io/core"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)

// Engine 数据库引擎
var Engine *xorm.Engine
var tables []names.TableName

func init() {
	tables = make([]names.TableName, 20)

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

func createTable(beans ...names.TableName) (err error) {
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
