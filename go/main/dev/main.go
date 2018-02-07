package main

import (
	"github.com/zaaksam/goman/go"
	"github.com/zaaksam/goman/go/config"
)

func main() {
	config.AppConf.Debug = true
	goman.Run("")
}
