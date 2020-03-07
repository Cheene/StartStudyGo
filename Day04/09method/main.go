package main

import "fmt"

//自定义类型添加方法 {基于别人的包李的类型创造自己的方法}
type myInt int

//不能给别的包类型变量添加方法，只能给自己的包的类型添加方法
func (m myInt) hello() {
	fmt.Println("Hello int")
}

func main() {
	m := myInt(100) //强制类型转换
	m.hello()
}
