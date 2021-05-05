package db

import (
	"github.com/zaaksam/goman/go/model"
	"xorm.io/xorm"
)

// DA .
type DA struct {
	*xorm.Session

	hasOrder bool
	groupBy  string
}

// NewDA 创建新的 DA
func NewDA() *DA {
	return &DA{
		Session: Engine.NewSession(),
	}
}

// Desc 降序
func (da *DA) Desc(colNames ...string) *xorm.Session {
	da.hasOrder = true
	return da.Session.Desc(colNames...)
}

// Asc 升序
func (da *DA) Asc(colNames ...string) *xorm.Session {
	da.hasOrder = true
	return da.Session.Asc(colNames...)
}

// OrderBy 排序
func (da *DA) OrderBy(order string) *xorm.Session {
	da.hasOrder = true
	return da.Session.OrderBy(order)
}

// GroupBy 分组
func (da *DA) GroupBy(keys string) *xorm.Session {
	da.groupBy = keys
	return da.Session.GroupBy(keys)
}

// GetList 获取分页数据
func (da *DA) GetList(md IModel, pageIndex, pageSize int) (list *model.ListModel, err error) {
	list = &model.ListModel{
		PageIndex: pageIndex,
		PageSize:  pageSize,
		Items:     md.NewItems(),
	}

	//克隆查询对象
	sessionCount := Engine.NewSession()
	defer sessionCount.Close()

	cond := da.Conds()
	if da.groupBy != "" {
		list.Total, err = sessionCount.Select("count(DISTINCT " + da.groupBy + ")").Where(cond).Count(md)
	} else {
		list.Total, err = sessionCount.Where(cond).Count(md)
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
	da.Limit(list.PageSize, recordIndex)

	err = da.Table(md).Find(list.Items)
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
