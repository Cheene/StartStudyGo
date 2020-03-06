package main

import (
	"fmt"
)

// 自定义类型和类型别名

//type后面跟的是类型 主要用域添加方法使用
type myInt int // 自定义类型

type yourInt = int

func main() {

	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 200
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	// rune 表示字符，底层是 int32
	var c rune
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
