// +build linux

package goman

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/zaaksam/goman/go/config"
)

func run(mode string) {
	logs.Info("====== 欢迎使用 " + config.AppConf.Name + " " + config.AppConf.Version + " ======")
	beego.Run()
}
