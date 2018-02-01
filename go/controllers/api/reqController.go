package api

import (
	"encoding/json"
	"errors"

	"github.com/zaaksam/goman/go/model"
	"github.com/zaaksam/goman/go/services"
)

type ReqController struct {
	BaseController
}

func (c *ReqController) Post() {
	body := c.GetString("req")
	if body == "" {
		c.SetError("req参数不能为空")
		return
	}

	req := &model.RequestModel{}
	err := json.Unmarshal([]byte(body), req)
	if err != nil {
		err = errors.New("请求参数反序列化错误：" + err.Error())
	} else if req.Method == "" {
		err = errors.New("method不能为空")
	} else if req.URL == "" {
		err = errors.New("url不能为空")
	}
	if err != nil {
		c.SetError(err)
		return
	}

	if req.N <= 0 {
		req.N = 1
	}

	if req.C <= 0 {
		req.C = 1
	}

	list, err := services.Req.Do(req)
	if err != nil {
		c.SetError(err)
		return
	}

	c.SetModel(list)
}
