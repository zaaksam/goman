package goman

import (
	"runtime"

	"github.com/astaxie/beego"
	// _ "github.com/mattn/go-sqlite3"
	"github.com/zaaksam/goman/go/config"
	// _ "github.com/zaaksam/goman/go/db"
	_ "github.com/zaaksam/goman/go/routers"
)

// Run 运行goman
func Run(mode string) {
	runtime.GOMAXPROCS(-1)

	beego.BConfig.AppName = config.AppConf.Name
	beego.BConfig.ServerName = config.AppConf.Name
	beego.BConfig.WebConfig.AutoRender = false
	beego.BConfig.WebConfig.DirectoryIndex = false
	beego.BConfig.CopyRequestBody = true

	beego.BConfig.Listen.HTTPAddr = config.AppConf.IP
	beego.BConfig.Listen.HTTPPort = config.AppConf.Port

	if config.AppConf.Debug {
		beego.BConfig.RunMode = beego.DEV

		// db.Engine.ShowSQL(true)
	} else {
		beego.BConfig.RunMode = beego.PROD
	}

	run(mode)
}
