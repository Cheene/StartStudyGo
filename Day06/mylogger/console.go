package mylogger

import (
	"fmt"
	"time"
)

//自定义一个日志库
// enable 返回bool 实现 l.level
func (l Logger) enable(level LogLevel) bool {
	return l.Level <= level
}

func log(lv LogLevel, msg string) {
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)

	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
}

func (l Logger) Debug(msg string) {
	if l.enable(DEBUG) {
		log(DEBUG, msg)
	}
}

func (l Logger) Trace(msg string) {
	if l.enable(TRACE) {
		log(TRACE, msg)
	}
}

func (l Logger) Error(msg string) {
	if l.enable(ERROR) {
		log(ERROR, msg)
	}
}

func (l Logger) Warning(msg string) {
	if l.enable(WARNING) {
		log(WARNING, msg)
	}
}

func (l Logger) Info(msg string) {
	if l.enable(INFO) {
		log(INFO, msg)
	}
}

func (l Logger) Fatal(msg string) {
	if l.enable(FATAL) {
		log(FATAL, msg)
	}
}
