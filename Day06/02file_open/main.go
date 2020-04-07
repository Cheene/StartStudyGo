package main

import (
	"fmt"
	"io"
	"os"
)

func f1() {
	fileObj, err := os.Open("./split.go")
	//defer fileObj.Close() // 当 error有数值的时候，fileObj 就是 nil
	// 下面遇到return的时候，会继续执行 fileObj.Close()
	if err != nil {
		fmt.Println("Error ", err)
		return
	}
	defer fileObj.Close()
}
func f2() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./aa.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer fileObj.Close()
	fileObj.Seek(3, 0) //光标开始移动
	s := []byte{'c', 'd', 'e'}
	fileObj.Write(s)
	//把剩余部分的内容写入到临时文件

	//删除源文件

	//将临时文件重新命名

}
func f3() {
	//打开要操作的文件
	fileObj, err := os.OpenFile("./aa.txt", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	//由于无法在文件中插入内容，所以需要先创建一个临时文件
	tempFile, err := os.OpenFile("aa.tmp", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644)

	if err != nil {
		fmt.Printf("OPEN FILE FAILED ,ERR:%v\n", err)
		return
	}
	defer tempFile.Close()
	//将原文件的内容的前面部分写入到临时文件
	//读--》 存入到 ret 里面
	var ret [1]byte
	n, err := fileObj.Read(ret[:])
	if err != nil {
		fmt.Printf("read from file failed:%v", err)
		return
	}
	//开始写
	tempFile.Write(ret[:n])
	fmt.Println("写入临时文件成功")
	//再将待添加的内容写入到临时文件
	tempAdd := []byte{'1', '2', '3', '4'}
	tempFile.Write(tempAdd)
	//把剩余部分的内容写入到临时文件
	//1光标移动到新的位置；
	fileObj.Seek(1, 0)
	//2把剩余的内容Copy过来
	var x [1024]byte
	for {
		n, err := fileObj.Read(x[:])
		if err == io.EOF {
			tempFile.Write(x[:n])
			break
		}
		if err != nil {
			fmt.Println("Error :%v", err)
			return
		}
		tempFile.Write(x[:n])
	}
	//删除源文件
	fileObj.Close()
	tempFile.Close()
	os.Rename("./aa.tmp", "./aa.txt")
	//将临时文件重新命名
}
func main() {
	f3()

}
