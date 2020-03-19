package mylogger

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

//定义级别的类型，基于内置的类型造一个内置的类型
type LogLevel uint8

//定义日志的级别
const (
	UNKNOW LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := fmt.Errorf("日志级别未知")
		return UNKNOW, err
	}
}

// Logger 日志对象
//日志的结构体
type Logger interface {
	Debug(format string, a ...interface{})
	Trace(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Info(format string, a ...interface{})
	Fatal(format string, a ...interface{})
	Error(format string, a ...interface{})
}

// NewLog 构造函数
// 1 字符串转level；2 构造一个Logger

func getInfo(n int) (funcName, fileName string, lineNo int) {
	//file 调用这个函数的是谁，line是行数
	pc, file, line, ok := runtime.Caller(n)

	if !ok {
		fmt.Println("runtime.Caller() failed ")
		return
	}
	//获取函数名
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	//获取文件的名称
	fileName = path.Base(file)
	//获取行号
	lineNo = line
	return
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case FATAL:
		return "FATAL"
	case INFO:
		return "INFO"
	case ERROR:
		return "ERROR"
	case WARNING:
		return "WARNING"
	default:
		return "UNKNOW"
	}
}
