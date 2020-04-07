package main

import (
	"fmt"
	"path"
	"runtime"
)

func f1() {
	//file 调用这个函数的是谁，line是行数
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime.Caller() failed ")
		return
	}
	//获取函数名
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)
	//获取文件的名称
	fmt.Println(path.Base(file)) // split.go
	fmt.Println(file)            //D:/StartStudyGo/Day06/06runtime_demo/split.go
	fmt.Println(line)
}
func f() {
	f1()
}
func main() {
	//file 调用这个函数的是谁，line是行数
	_, file, line, ok := runtime.Caller(1)

	if !ok {
		fmt.Println("runtime.Caller() failed ")
	}

	fmt.Println(file)

	fmt.Println(line)

	f()
}
