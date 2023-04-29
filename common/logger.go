package common

import (
	"fmt"
	log2 "log"
	"runtime"

	"github.com/mgcicd/cicd-core/logs"
)

const logtype = "envoy-xds"

func Info(logger string, message string, mesc int64) {
	go logs.NewKafkaLogger().Info(logtype, logger, message)
}

func Warn(logger string, exception string, message string) {
	ConsoleLog("logger:" + logger)
	ConsoleLog("exception:" + exception)
	ConsoleLog("message:" + message)
	go logs.NewKafkaLogger().Warn(logtype, logger, message)
}

func Error(logger string, exception string, message string) {
	ConsoleLog("logger:" + logger)
	ConsoleLog("exception:" + exception)
	ConsoleLog("message:" + message)
	go logs.NewKafkaLogger().Error(logtype, logger, message)
}

func Debug(logger string, exception string, message string) {
	go logs.NewKafkaLogger().Debug(logtype, logger, message)
}

func ConsoleLog(v ...interface{}) {
	log2.SetFlags(log2.LstdFlags | log2.Lshortfile)
	log2.Println(v)
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}
