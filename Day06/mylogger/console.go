package mylogger

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	Level LogLevel
}

func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

//自定义一个日志库
// enable 返回bool 实现 l.level
func (l ConsoleLogger) enable(level LogLevel) bool {
	return l.Level <= level // l:Trace; level:DEBUG; 显然不能打印 DEBUG 类型
}

func log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...) //拼接 format 和 参数
	fmt.Println(msg)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)

	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

//Debug 仅能由Logger结构体调用
func (l ConsoleLogger) Debug(format string, a ...interface{}) {
	if l.enable(DEBUG) {
		log(DEBUG, format, a)
	}
}

func (l ConsoleLogger) Trace(format string, a ...interface{}) {
	if l.enable(TRACE) {
		log(TRACE, format, a)
	}
}

func (l ConsoleLogger) Error(format string, a ...interface{}) {
	if l.enable(ERROR) {
		log(ERROR, format, a)
	}
}

func (l ConsoleLogger) Warning(format string, a ...interface{}) {
	if l.enable(WARNING) {
		log(WARNING, format, a)
	}
}

func (l ConsoleLogger) Info(format string, a ...interface{}) {
	if l.enable(INFO) {
		log(INFO, format, a)
	}
}

func (l ConsoleLogger) Fatal(format string, a ...interface{}) {
	if l.enable(FATAL) {
		log(FATAL, format, a)
	}
}
