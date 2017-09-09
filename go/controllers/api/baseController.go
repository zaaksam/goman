package api

import (
	"github.com/astaxie/beego"
)

// BaseController API基础控制器
type BaseController struct {
	beego.Controller
	code  ResponseCodeType
	msg   string
	data  map[string]interface{}
	model interface{}
}

// ResponseCodeType 响应代码类型
type ResponseCodeType int

const (
	// ResponseOK 正常响应
	ResponseOK ResponseCodeType = 10000
	// ResponseError 错误响应
	ResponseError ResponseCodeType = 90000
	// ResponseUnauthorized 没有权限
	ResponseUnauthorized ResponseCodeType = 90100
)

// Prepare beego Controller Prepare事件
func (c *BaseController) Prepare() {
	c.code = ResponseOK
}

// SetModel 设置响应的数据模型
func (c *BaseController) SetModel(model interface{}) {
	c.model = model
}

// SetData 设置响应数据
func (c *BaseController) SetData(key string, val interface{}) {
	if c.data == nil {
		c.data = make(map[string]interface{})
	}
	c.data[key] = val
}

// SetError 设置错误响应信息
func (c *BaseController) SetError(options ...interface{}) {
	c.code = ResponseError

	for i, l := 0, len(options); i < l; i++ {
		switch opt := options[i].(type) {
		case ResponseCodeType:
			c.code = opt
		case string:
			c.msg = opt
		case error:
			c.msg = opt.Error()
		}
	}
}

// Finish beego Controller Finish事件
func (c *BaseController) Finish() {
	m := make(map[string]interface{})
	m["code"] = c.code
	m["msg"] = c.msg

	if c.model != nil {
		m["data"] = c.model
	} else {
		if len(c.data) == 0 {
			m["data"] = struct{}{}
		} else {
			m["data"] = c.data
		}
	}

	c.Data["json"] = m
	c.ServeJSON()
}
