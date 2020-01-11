package main

import "fmt"

//常量定义后不能改变
const pi = 3.14159265382917

//批量声明常量
const (
	statusOK = 200
	notFound = 404
)

//默认 n1,n2,n3 的数值相同
const (
	n1 = 100
	n2
	n3
)

// iota 常量计数器
// 1 const 出现的时候，iota 初始化为0
// 2 const 内部的时候，每增加一行，iota 自增+1
const (
	c1 = iota
	c2 = iota
)
const (
	c3 = iota
)

//多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2

	d3, d4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println(pi)
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println("c1", c1)
	fmt.Println("c2", c2)
	fmt.Println("c3", c3)

	fmt.Println("d1", d1)
	fmt.Println("d2", d2)
	fmt.Println("d3", d3)
	fmt.Println("d4", d4)

}
