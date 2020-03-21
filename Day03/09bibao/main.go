package main

import "fmt"

// 1 f1 里面只有一个参数；即 func()类型
// 2 在 f1 里面执行函数 ff3
// 3 把 ff3 的结果返回

//闭包,
//1 类型不匹配的函数
func f1(f func()) {
	fmt.Println("This is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("This is f2")
	fmt.Println(x + y)
}

//定义一个函数对f2 进行包装
func f3(f func(int, int) int) func() { // ① 参数的类型是 f2 ， 返回值的类型是 f1参数所需要的类型
	tmp := func() {
		fmt.Println("This is Tmp")
	}
	return tmp
}

func lixiang(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}

//要求

func ff3(x, y int, f func(int, int)) func() {
	return func() {
		f(x, y)
	}
}

func main() {
	f1(ff3(200, 10000, f2))
	//把原来需要传递两个int类型的参数包装成一个不需要传递参数的类型
	//ret := lixiang(f2,3,522)
	//ret()
}
