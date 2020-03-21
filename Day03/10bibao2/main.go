package main

import "fmt"

//闭包就是一个函数，这个函数的特点是包含了他外部作用域的一个变量

func adder() func(int) int {
	x := 100
	return func(y int) int {
		x += y
		return x
	}
}
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	//为何不是 func(int)int ?
	ret := adder() // ret 是函数在内存的首地址
	fmt.Printf("%T\n", ret)
	ret2 := ret(200)  // 能够调用函数体外部的 x，除了返回函数，还返回变量 x
	fmt.Println(ret2) // x+y = 300
	//
	add2 := adder2(2)
	fmt.Println(add2(2))
}
