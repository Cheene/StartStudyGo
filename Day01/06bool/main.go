package main

import "fmt"

func main() {
	//布尔值
	//1 默认值为 false; 2 不允许将整型强制转换为布尔类型
	//3 布尔类型属于独立的类型，也无法直接转换为整型或参与整型的计算。
	b1 := true
	var b2 bool
	fmt.Printf("%v\n", b1)
	fmt.Printf("%v\n", b2)
}
