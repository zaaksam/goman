package routers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"
	"github.com/rakyll/statik/fs"
	"github.com/zaaksam/goman/go/config"
	"github.com/zaaksam/goman/go/controllers"
	"github.com/zaaksam/goman/go/controllers/api"
	_ "github.com/zaaksam/goman/go/statik"
)

func init() {
	beego.ErrorController(&controllers.ErrController{})

	api := beego.NewNamespace("/api",
		beego.NSRouter("/app", &api.AppController{}),
		beego.NSRouter("/req", &api.ReqController{}),
		beego.NSRouter("/log", &api.LogController{}),
	)

	web := beego.NewNamespace("/web",
		beego.NSRouter("/config.js", &controllers.WebController{}, "get:Config"),
		beego.NSRouter("/welcome", &controllers.WebController{}, "get:Welcome"),
		beego.NSRouter("/*", &controllers.WebController{}),
	)

	beego.Get("/", func(ctx *context.Context) {
		ctx.Redirect(302, "/web")
	})

	var staticHandler http.Handler
	if config.AppConf.Debug {
		// dir := path.Dir(os.Args[0])
		dir, _ := os.Getwd()
		dir = filepath.ToSlash(dir) // .../goman/go/main/xxx
		dirs := strings.Split(dir, "/")
		dir = strings.Join(dirs[0:len(dirs)-3], "/") // .../goman
		staticHandler = http.StripPrefix("/static/", http.FileServer(http.Dir(dir+"/web/static")))
	} else {
		statikFS, err := fs.New()
		if err != nil {
			logs.Error("statik获取失败：", err)
		}
		staticHandler = http.StripPrefix("/static/", http.FileServer(statikFS))
	}

	static := beego.NewNamespace("/static", beego.NSGet("/*", func(ctx *context.Context) {
		staticHandler.ServeHTTP(ctx.ResponseWriter, ctx.Request)
	}))

	beego.AddNamespace(static, api, web)
}
