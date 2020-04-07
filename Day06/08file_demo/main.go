package main

import (
	"fmt"
	"os"
)

func main() {
	fileObj, err := os.Open("./split.go")

	if err != nil {
		fmt.Printf("ERR: %d\n", fileObj)
	}
	//1 文件对象的类型
	fmt.Printf("%T\n", fileObj)
	//2.文件的大小
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("ERR: %v\n", err)
	}
	fmt.Printf("%v\n", fileInfo.Size())
}
