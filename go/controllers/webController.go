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

// Get 主页
func (c *WebController) Get() {
	var unix int64
	if config.AppConf.Debug {
		unix = time.Now().Unix()
	} else {
		unix = config.AppConf.Started
	}
	unixStr := "?t=" + strconv.FormatInt(unix, 10)

	body := strings.Replace(views.Index, "{{.unix}}", unixStr, -1)
	body = strings.Replace(body, "{{.appName}}", config.AppConf.Name, -1)

	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	c.Ctx.Output.Body([]byte(body))
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
		UserAgent string `json:"userAgent"`
	}{
		Name:      config.AppConf.Name,
		Version:   config.AppConf.Version,
		UserAgent: c.Ctx.Input.Header("User-Agent"),
	}

	j, _ := json.Marshal(conf)
	body := fmt.Sprintf("var C=%s;", string(j))

	c.Ctx.Output.Body([]byte(body))
}
