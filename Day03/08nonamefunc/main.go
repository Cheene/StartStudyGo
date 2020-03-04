package main

import "fmt"

//匿名函数，常用于函数的内部
var f1 = func(x, y int) {
	fmt.Println(x + y)
}

func main() {
	f1 := func() {
		fmt.Println("Inner Func")
	}
	f1()
	//立即执行函数
	func() {
		fmt.Println("Now")
	}()
	func(x, y int) {
		fmt.Println(x + y)
	}(100, 30)
}
