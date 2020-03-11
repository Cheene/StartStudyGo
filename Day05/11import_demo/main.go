package main

import (
	"../10calc"
	"fmt"
)

//main 只有 main 包可以形成可执行文件

var x = 100

func init() {
	fmt.Println("自动执行")
	fmt.Println(x)
}

func main() {
	var a = calc.Add(1, 2)
	fmt.Println(a)
}
