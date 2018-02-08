// +build !linux

package goman

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/zaaksam/goman/go/config"
	"github.com/zaaksam/goman/go/services"
	"github.com/zserge/webview"
)

func run(mode string) {
	if mode == "web" {
		config.AppConf.RunMode = config.RUN_MODE_WEB
	} else if mode == "app" {
		config.AppConf.RunMode = config.RUN_MODE_APP
	}

	url := fmt.Sprintf("%s/web", config.AppConf.Host)

	if config.AppConf.OS == config.OS_TYPE_LINUX {
		logs.Info("====== 欢迎使用 " + config.AppConf.Name + " " + config.AppConf.Version + "（Web模式） ，关闭此窗口即可退出程序 ======")

		if !config.AppConf.Debug {
			time.Sleep(1 * time.Second)
			go openBrowser(url)
		}

		beego.Run()

		return
	}

	go beego.Run()

	settings := webview.Settings{
		Title:     config.AppConf.Name,
		URL:       url,
		Width:     800,
		Height:    600,
		Resizable: true,
		Debug:     config.AppConf.Debug,
		ExternalInvokeCallback: func(w webview.WebView, data string) {
			switch data {
			case "check":
				isConnect := services.App.IsConnect()
				if isConnect {
					callJs(w, "connect")
				} else {
					callJs(w, "disconnect")
				}
			case "open":
				openBrowser(url)
			case "github":
				openBrowser("https://github.com/zaaksam/goman")
			case "gitee":
				openBrowser("https://gitee.com/zaaksam/goman")
			case "close":
				w.Terminate()
				w.Exit()
			}

		},
	}

	if config.AppConf.RunMode == config.RUN_MODE_WEB {
		settings.URL += "/welcome"
		settings.Width = 500
		settings.Height = 300
		settings.Resizable = false
	}

	w := webview.New(settings)
	w.Run()
}

func callJs(w webview.WebView, data string) {
	err := w.Eval("vm.callJs('" + data + "');")
	if err != nil {
		logs.Error("向js发送消息失败：", err)
	}
}

func openBrowser(url string) {
	cmds := map[config.OSType][]string{
		config.OS_TYPE_WINDOWS: []string{"cmd", "/c", "start", ""},
		config.OS_TYPE_DARWIN:  []string{"open", ""},
		config.OS_TYPE_LINUX:   []string{"xdg-open", ""},
	}

	args, ok := cmds[config.AppConf.OS]
	if !ok {
		logs.Error("无法打开浏览器：", config.AppConf.OS)
		return
	}

	args[len(args)-1] = url

	c := exec.Command(args[0], args[1:]...)
	err := c.Start()
	if err != nil {
		logs.Critical("打开浏览器失败：", err)
	}
}
