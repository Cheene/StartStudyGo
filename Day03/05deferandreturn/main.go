package main

import "fmt"

// 一 返回值赋值
// 二 defer
// 三 真正的RET返回
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //返回值是 5
}

func f2() (x int) {
	defer func() {
		x++ // 修改的就是 x
	}()
	return 5 // 返回值是 x
}

func f3() (y int) {
	x := 5 // 第一步 返回值 y = x = 5
	defer func() {
		x++ // 第二步，修改x
	}()
	return x // 返回 y的值
}
func f4() (x int) {
	defer func(x int) {
		x++ // x 是函数内部的副本
	}(x) //函数传参传的是函数的副本
	return 5
}

func f6() (x int) {
	defer func(x *int) {
		*x++
	}(&x) //函数传参传的是函数的副本
	return 5
}

func main() {
	fmt.Println(f1()) // 5, x = 6
	fmt.Println(f2()) // 重点
	fmt.Println(f3()) // 5, x = 6
	fmt.Println(f4()) // 5, x = 6
	fmt.Println(f6()) // 5, x = 6

}
