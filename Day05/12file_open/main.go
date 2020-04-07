package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//打开文件
func readFormFile1(file os.File) {

	//读文件
	//你想读到哪里
	n := 32
	var tmp = make([]byte, n)

	for {

		n1, err := file.Read(tmp) //一次读取128个byte
		if err != nil {
			fmt.Println("Error %v", err)
			return
		}
		fmt.Printf("读了%d个字节", n1)
		fmt.Println(string(tmp[:n]))
		if n1 < n { // 如果n < 最大数说明已经全部读完
			return
		}
	}
}

func readFormFileBufio(file *os.File) {
	//先创建一个用来读内容的对象
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Failed %v", err)
			return
		}
		fmt.Print(line)
	}
}
func readFromFileByIOUtil() {

	ret, err := ioutil.ReadFile("./split.go")
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(ret)

}

func main() {
	file, err := os.Open("./111.txt")

	if err != nil {
		fmt.Println("filed open %v", err)
	}
	defer file.Close() // 防止忘记关闭文件
	readFormFile1(*file)
	readFormFileBufio(file)
	readFromFileByIOUtil()
}
