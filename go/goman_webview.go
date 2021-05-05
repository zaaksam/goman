// +build !linux

package goman

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/webview/webview"
	"github.com/zaaksam/goman/go/config"
	"github.com/zaaksam/goman/go/services"
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

	// Resizable: true
	w := webview.New(config.AppConf.Debug)
	w.SetTitle(config.AppConf.Name)

	if config.AppConf.RunMode == config.RUN_MODE_WEB {
		// Resizable: false
		w.SetSize(500, 300, webview.HintFixed)
		w.Navigate(url + "/welcome")
	} else {
		// Resizable: true
		w.SetSize(800, 600, webview.HintNone)
		w.Navigate(url)
	}

	w.Bind("check", func() {
		isConnect := services.App.IsConnect()
		if isConnect {
			callJs(w, "connect")
		} else {
			callJs(w, "disconnect")
		}
	})

	w.Bind("open", func() {
		openBrowser(url)
	})

	w.Bind("github", func() {
		openBrowser("https://github.com/zaaksam/goman")
	})

	w.Bind("gitee", func() {
		openBrowser("https://gitee.com/zaaksam/goman")
	})

	w.Bind("close", func() {
		w.Terminate()
		w.Destroy()
	})

	w.Run()
}

func callJs(w webview.WebView, data string) {
	w.Eval("vm.callJs('" + data + "');")
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
