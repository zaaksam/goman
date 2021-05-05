package config

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"runtime"
	"strings"
	"time"
)

// AppConf 应用全局参数
var AppConf appConf

func init() {
	AppConf.goos = runtime.GOOS
	AppConf.goarch = runtime.GOARCH

	if AppConf.goos == "windows" {
		AppConf.OS = OS_TYPE_WINDOWS
	} else if AppConf.goos == "darwin" {
		AppConf.OS = OS_TYPE_DARWIN
	} else if AppConf.goos == "linux" {
		AppConf.OS = OS_TYPE_LINUX
	}

	if AppConf.OS == OS_TYPE_UNSUPPORT {
		panic("不支持 " + AppConf.goos + " 系统的编译运行")
	}

	if AppConf.goarch == "amd64" {
		AppConf.Arch = ARCH_TYPE_AMD64
	}

	if AppConf.Arch == ARCH_TYPE_UNSUPPORT {
		panic("不支持 " + AppConf.goarch + " 核心的编译运行")
	}

	mode := ""
	flag.BoolVar(&AppConf.Debug, "debug", false, "调试模式，默认：false")
	flag.StringVar(&AppConf.IP, "ip", "", "监听的IP地址，默认：127.0.0.1")
	flag.IntVar(&AppConf.Port, "port", 0, "服务端口，默认：随机")
	flag.StringVar(&mode, "mode", "app", "运行模式：web：Web模式；app：App模式(试验)；默认：app")
	flag.Parse()

	if AppConf.IP == "" {
		AppConf.IP = "127.0.0.1"
	}

	if AppConf.Port <= 0 {
		AppConf.Port = newPort(AppConf.IP)
	}

	AppConf.Host = fmt.Sprintf("http://%s:%d", AppConf.IP, AppConf.Port)

	switch strings.ToLower(mode) {
	case "web":
		AppConf.RunMode = RUN_MODE_WEB
	case "app":
		AppConf.RunMode = RUN_MODE_APP
	default:
		AppConf.RunMode = RUN_MODE_WEB
	}

	AppConf.Name = "goman"
	AppConf.Version = "0.4.0"
	AppConf.Started = time.Now().Unix()
}

type appConf struct {
	goos    string
	goarch  string
	OS      OSType
	Arch    ArchType
	Name    string
	Version string
	Debug   bool
	Host    string
	IP      string
	Port    int
	Started int64
	RunMode RunModeType
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
