package main

import "fmt"

//推荐驼峰式命名 studnetName
var (
	s    string
	i    int
	isOK bool
)

func foo() (int, string) {
	return 2, "Chenene"
}

func main() {
	s = "XX"
	i = 16
	isOK = true
	//% 是占位符
	fmt.Printf("s:%s,i:%d,isOK:%v\n", s, i, isOK)
	//声明变量同时赋值
	var userName string = "chenene"
	fmt.Println(userName)
	// 类型推导，根据值来判断变量属于什么类型
	var s11 = "29"
	fmt.Println(s11)

	// 短变量声明 仅能够在函数变量声明中使用
	t := 1
	fmt.Printf("t:%d", t)
	// 匿名变量，需要忽略某个值的时候， 无内存，不分配空间

	x, _ := foo()

	_, y := foo()
	fmt.Println()
	fmt.Println(x)
	fmt.Println(y)
}
