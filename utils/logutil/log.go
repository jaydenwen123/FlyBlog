package logutil

import (
	"github.com/astaxie/beego/logs"
	"path/filepath"
	"runtime"
)

type LogInfo struct {
	FileName   string
	MethodName string
	LogMsg     string
}

func Debug(info LogInfo) {
	logs.Debug("executing the file [", info.FileName, "] method <", info.MethodName, ">\t", "info:", info.LogMsg)
}

func Info(info LogInfo) {
	logs.Info("executing the file [", info.FileName, "] method <", info.MethodName, ">\t", "info:", info.LogMsg)
}
func Error(info LogInfo) {
	logs.Error("executing the file [", info.FileName, "] method <", info.MethodName, ">\t", "info:", info.LogMsg)
}
func Warn(info LogInfo) {
	logs.Warn("executing the file [", info.FileName, "] method <", info.MethodName, ">\t", "info:", info.LogMsg)
}
func Fatal(info LogInfo) {
	logs.Trace("executing the file [", info.FileName, "] method <", info.MethodName, ">\t", "info:", info.LogMsg)
}

func CurrentFileName() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		logs.Error("system error")
	}
	//_,file=filepath.Split(file)
	return filepath.Base(file)
}
