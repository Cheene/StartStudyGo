package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//向文件中填写日志
type FileLogger struct {
	Level         LogLevel
	filePath      string //日志文件路径
	fileName      string //日志文件的名字
	errorFileName string //错误日志的文件
	maxFileSize   uint64 //最大文件的大小
	fileObj       *os.File
	errFileObj    *os.File
}

func NewFileLogger(level, filepath, filename string, mfs uint64) *FileLogger {
	logLevel, err := parseLogLevel(level)
	if err != nil {
		fmt.Println("This is Error: ", err)
		panic(err) // 不能直接return
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    filepath,
		fileName:    filename,
		maxFileSize: mfs,
	}
	err = fl.initFIle()
	if err != nil {
		panic(err)
	}
	return fl
}

//根据新建的形成格式，这里如果不是指针格式的话，传递的只能是副本，即原来的值还是 nil
func (f *FileLogger) initFIle() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	//打开日志文件
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("oprn log file failed,err: ", err)
		return err
	}
	//打开错误日志文件
	fileErrorObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("oprn log file failed,err: ", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = fileErrorObj
	return err
}

func (f FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...) //拼接 format 和 参数
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		//如果要记录的日志的级别大于或等于ERROR，便在 ERROR 文件中再记录一遍
		if lv >= ERROR {
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) enable(level LogLevel) bool {
	return f.Level <= level // l:Trace; level:DEBUG; 显然不能打印 DEBUG 类型
}

//Debug 仅能由Logger结构体调用
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLogger) Trace(format string, a ...interface{}) {
	f.log(TRACE, format, a...)
}

func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
