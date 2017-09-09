package config

import (
	"flag"
	"strings"
	"time"
)

// AppConf 应用全局参数
var AppConf appConf

func init() {
	mode := ""
	flag.BoolVar(&AppConf.Debug, "debug", false, "调试模式，默认：false")
	flag.StringVar(&AppConf.IP, "ip", "", "监听的IP地址，默认：127.0.0.1")
	flag.IntVar(&AppConf.Port, "port", 0, "服务端口，默认：8080")
	flag.StringVar(&mode, "mode", "web", "运行模式：web：Web模式；app：App模式(暂不支持)；默认：web")
	flag.Parse()

	if AppConf.IP == "" {
		AppConf.IP = "127.0.0.1"
	}

	if AppConf.Port <= 0 {
		AppConf.Port = 8080
	}

	switch strings.ToLower(mode) {
	case "web":
		AppConf.IsWebMode = true
	// case "app":
	// 	AppConf.IsAppMode = true
	default:
		AppConf.IsWebMode = true
	}

	AppConf.Name = "goman"
	AppConf.Version = "0.1.0"
	AppConf.Started = time.Now().Unix()
}

type appConf struct {
	Name      string
	Version   string
	Debug     bool
	IP        string
	Port      int
	Started   int64
	IsAppMode bool
	IsWebMode bool
}
