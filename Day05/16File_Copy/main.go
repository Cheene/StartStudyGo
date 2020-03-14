package main

import (
	"fmt"
	"io"
	"os"
)

/**
1 打开准备读的文件
2 打开或者创建文件，并清空内容
3 将读的文件内容转移到写的文件里面
*/
func CopyFile(dstNmae, scrName string) (written int64, err error) {
	//打开准备要读的文件
	src, err := os.Open(scrName) // Open
	if err != nil {
		fmt.Printf("open %s failed%v\n", src, err)
	}
	defer src.Close()
	//以写|创建的方式打开目标文件
	dest, err := os.OpenFile(dstNmae, os.O_WRONLY|os.O_CREATE, 0644) //OpenFile
	if err != nil {
		fmt.Printf("Open %s failed,%v\n", dest, err)
	}
	defer dest.Close()
	return io.Copy(dest, src)
}

func main() {
	_, err := CopyFile("./desc.txt", "./111.txt")
	if err != nil {
		fmt.Println("copy file failed,err :", err)
		return
	}
	fmt.Println("Copy done!")
}
