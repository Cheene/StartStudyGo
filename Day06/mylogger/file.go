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

//如果文件大小大于最大值，就要返回 true
func (f *FileLogger) checkSize(file *os.File) bool {
	fileinfo, err := file.Stat()

	if err != nil {
		fmt.Printf("fileInfo is  %v", err)
		return false
	}
	return uint64(fileinfo.Size()) >= f.maxFileSize
}
func (f *FileLogger) splitLogFile(file *os.File) (*os.File, error) {
	//获取文件名
	fileInfo, err := file.Stat()

	if err != nil {
		fmt.Printf("get file info failed,%v ", err)
		return nil, err
	}
	now := time.Now().Format("20060102150405000")
	logName := path.Join(f.filePath, fileInfo.Name())   //获取日志完整路径
	logNewName := fmt.Sprintf("%s.bak%s", logName, now) //拼接一个日志备份的名字
	//1.关闭当前的日志文件
	file.Close()
	//2.备份当前的文件 rename
	os.Rename(logName, logNewName)
	//3.打开一个新的文件对象并重新赋值给obj
	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error Open File: %v\n", err)
		return nil, err
	}
	return fileObj, err
}
func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...) //拼接 format 和 参数
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		//需要进行切割日志文件
		if f.checkSize(f.fileObj) {
			newFile, err := f.splitLogFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), getLogString(lv), fileName, funcName, lineNo, msg)
		//如果要记录的日志的级别大于或等于ERROR，便在 ERROR 文件中再记录一遍
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newErrorFile, err := f.splitLogFile(f.errFileObj)
				if err != nil {
					return
				}
				f.errFileObj = newErrorFile
			}
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
