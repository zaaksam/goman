package logger

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"
	// "github.com/zaaksam/goman/go/db"
)

// Info 普通级别日志
func Info(v ...interface{}) {
	insert("Info", v...)
}

// Debug 调试级别日志
func Debug(v ...interface{}) {
	insert("Debug", v...)
}

// Error 错误级别日志
func Error(v ...interface{}) {
	insert("Error", v...)
}

/// Warning 警告级别日志
func Warning(v ...interface{}) {
	insert("Warning", v...)
}

// Critical 严重级别日志
func Critical(v ...interface{}) {
	insert("Critical", v...)
}

func insert(typ string, v ...interface{}) {
	vLen := len(v)
	if vLen == 0 {
		return
	}

	content := fmt.Sprintf(strings.Repeat(" %v", vLen), v...)

	switch typ {
	case "Info":
		logs.Info(content)
	case "Debug":
		logs.Debug(content)
	case "Error":
		logs.Error(content)
	case "Warning":
		logs.Warning(content)
	case "Critical":
		logs.Critical(content)
	}

	// md := &model.LogModel{
	// 	Type:    typ,
	// 	Content: content,
	// 	Created: time.Now().Unix(),
	// }

	// _, err := db.Engine.Insert(md)
	// if err != nil {
	// 	j, _ := json.Marshal(md)
	// 	logs.Critical(fmt.Sprintf("日志 [%s] 写入log日志错误：", string(j)), err)
	// }
}
