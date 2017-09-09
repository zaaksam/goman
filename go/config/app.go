package config

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

// AppConf 应用全局参数
var AppConf appConf

func init() {
	mode := ""
	flag.BoolVar(&AppConf.Debug, "debug", false, "调试模式，默认：false")
	flag.StringVar(&AppConf.IP, "ip", "", "监听的IP地址，默认：127.0.0.1")
	flag.IntVar(&AppConf.Port, "port", 0, "服务端口，默认：随机")
	flag.StringVar(&mode, "mode", "web", "运行模式：web：Web模式；app：App模式(暂不支持)；默认：web")
	flag.Parse()

	if AppConf.IP == "" {
		AppConf.IP = "127.0.0.1"
	}

	if AppConf.Port <= 0 {
		AppConf.Port = newPort(AppConf.IP)
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
	AppConf.Version = "0.1.2"
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

//newPort 查找可用端口
func newPort(ip string) int {
	i := 0
	for {
		if i > 10 {
			return 8080
		}

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		port := r.Intn(60000)
		if port <= 0 {
			continue
		}

		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
		if err != nil {
			if strings.Contains(err.Error(), "refused") {
				return port
			}

			i++
			continue
		}
		conn.Close()
	}
}
