package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/zaaksam/goman/go/config"
	"github.com/zaaksam/goman/go/views"
)

// WebController 网页控制器
type WebController struct {
	beego.Controller
}

// Welcome 浏览器版本欢迎页
func (c *WebController) Welcome() {
	body := c.formatHTML(views.Welcome)

	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	c.Ctx.Output.Body([]byte(body))
}

// Get 主页
func (c *WebController) Get() {
	body := c.formatHTML(views.Index)

	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	c.Ctx.Output.Body([]byte(body))
}

func (c *WebController) formatHTML(body string) string {
	windowsDebug := ""

	var unix int64
	if config.AppConf.Debug {
		unix = time.Now().Unix()

		if config.AppConf.OS == config.OS_TYPE_WINDOWS {
			windowsDebug = `<script type="text/javascript" src="https://getfirebug.com/firebug-lite.js"></script>`
		}
	} else {
		unix = config.AppConf.Started
	}
	unixStr := "?t=" + strconv.FormatInt(unix, 10)

	body = strings.Replace(body, "{{.unix}}", unixStr, -1)
	body = strings.Replace(body, "{{.appName}}", config.AppConf.Name, -1)
	body = strings.Replace(body, "{{.windowsDebug}}", windowsDebug, -1)

	return body
}

// Config 配置文件
func (c *WebController) Config() {
	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "application/javascript; charset=utf-8")
	}
	c.Ctx.WriteString("")

	conf := &struct {
		Name      string `json:"appName"`
		Version   string `json:"appVersion"`
		RunMode   string `json:"runMode"`
		URL       string `json:"url"`
		UserAgent string `json:"userAgent"`
	}{
		Name:      config.AppConf.Name,
		Version:   config.AppConf.Version,
		URL:       fmt.Sprintf("http://%s:%d/web", config.AppConf.IP, config.AppConf.Port),
		UserAgent: c.Ctx.Input.Header("User-Agent"),
	}

	if config.AppConf.RunMode == config.RUN_MODE_WEB {
		conf.RunMode = "web"
	} else if config.AppConf.RunMode == config.RUN_MODE_APP {
		conf.RunMode = "app"
	} else if config.AppConf.RunMode == config.RUN_MODE_DOCKER {
		conf.RunMode = "docker"
	}

	j, _ := json.Marshal(conf)
	body := fmt.Sprintf("var C=%s;", string(j))

	c.Ctx.Output.Body([]byte(body))
}
