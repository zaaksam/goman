package api

import (
	"github.com/zaaksam/goman/go/services"
)

type AppController struct {
	BaseController
}

func (c *AppController) Get() {
	//心跳
	services.App.Heartbeat()
}
