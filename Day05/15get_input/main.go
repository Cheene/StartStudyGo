package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入的时候如果有空格

func useScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scanln(&s)
	fmt.Printf("%s\n", s)
}

func userBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println(s)
}
func main() {
	//useScan()
	userBufio()

}
