package main

import "fmt"

//函数类型
// func()
func f1() {
	fmt.Println("hello china")
}

//func() int
func f2() int {
	fmt.Println("hello f2")
	return 0
}

//函数也可以作为参数的类型
func f3(x func()) {
	x()
}

// f4 的形参是  func() int 类型
func f4(x func() int) {
	x()
}

//函数可以作为返回值
func f6(x, y int) int {
	return x + y
}
func f5(x func() int) func(int, int) int {
	ret := f6
	return ret
}

func main() {
	a := f1
	fmt.Printf("%T\n", a)
	b := f2
	fmt.Printf("%T\n", b)
	f3(f1)
	f4(f2)
	f7 := f5(f2)
	fmt.Println(f7) // 返回值是一个内存地址
}
