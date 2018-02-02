package goman

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	// _ "github.com/mattn/go-sqlite3"
	"github.com/zaaksam/goman/go/config"
	// _ "github.com/zaaksam/goman/go/db"
	_ "github.com/zaaksam/goman/go/routers"
	"github.com/zserge/webview"
)

// Run 运行goman
func Run(mode string) {
	if mode == "web" {
		config.AppConf.IsAppMode = false
		config.AppConf.IsWebMode = true
	} else if mode == "app" {
		config.AppConf.IsAppMode = true
		config.AppConf.IsWebMode = false
	}

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

	if config.AppConf.IsWebMode {
		logs.Info("====== 欢迎使用 " + config.AppConf.Name + " " + config.AppConf.Version + "（Web模式） ，关闭此窗口即可退出程序 ======")

		if !config.AppConf.Debug {
			go openBrowser()
		}

		beego.Run()
	} else if config.AppConf.IsAppMode {
		//logs.Info("====== 欢迎使用 " + config.AppConf.Name + " " + config.AppConf.Version + "（App模式） ，关闭此窗口即可退出程序 ======")

		go beego.Run()

		// w := webview.New(webview.Settings{
		// 	Title:     config.AppConf.Name,
		// 	URL:       fmt.Sprintf("http://%s:%d/web", config.AppConf.IP, config.AppConf.Port),
		// 	Width:     800,
		// 	Height:    600,
		// 	Resizable: true,
		// 	ExternalInvokeCallback: func(w webview.WebView, data string) {
		// 		logs.Info("收到JS信息：", data)
		// 	},
		// })
		// defer w.Exit()

		// w.Run()

		err := webview.Open(config.AppConf.Name, fmt.Sprintf("http://%s:%d/web", config.AppConf.IP, config.AppConf.Port), 1024, 768, true)
		if err != nil {
			logs.Critical("webview启动失败：", err)
		}
	}
}

func openBrowser() {
	time.Sleep(2 * time.Second)

	cmds := map[string][]string{
		"windows": []string{"cmd", "/c", "start", ""},
		"darwin":  []string{"open", ""},
		"linux":   []string{"xdg-open", ""},
	}

	osStr := runtime.GOOS

	args, ok := cmds[osStr]
	if !ok {
		logs.Error("非预设系统，无法打开浏览器：", osStr)
		return
	}

	args[len(args)-1] = fmt.Sprintf("http://%s:%d/web", config.AppConf.IP, config.AppConf.Port)

	c := exec.Command(args[0], args[1:]...)
	err := c.Start()
	if err != nil {
		logs.Critical("打开浏览器失败：", err)
	}
}
