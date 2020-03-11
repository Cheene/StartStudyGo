package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func writeDemo2() {
	file, err := os.OpenFile("./01.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("Err: %v", err)
		return
	}
	defer file.Close()
	//创建一个写的对象
	wr := bufio.NewWriter(file)
	wr.WriteString("Chenenee\n")
	wr.Flush()
}

//直接写文件
func writeDemo3() {
	str := "Hello 沙河"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)

	if err != nil {
		fmt.Println("ERR ", err)
	}
}

//打开文件
func main() {
	// 111 1- 创建 2-追加/清空 3- 其他
	file, err := os.OpenFile("./01.txt", os.O_APPEND|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		fmt.Printf("Err: %v", err)
		return
	}
	//write 字节
	file.Write([]byte("zhoulinmengbbile\n")) // 字符串转字节切片
	//writeString 字符串
	file.WriteString("Chenene hhhh哈哈哈\n")

	writeDemo2()
	writeDemo3()
}
