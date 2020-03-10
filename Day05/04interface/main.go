package main

import "fmt"

//接口示例2
//不管什么struct,只要是调用run()方法就可以
type car interface {
	run()
}

type falali struct {
	brand string
}

type dazhong struct {
	brand string
}

//driver 函数接收一个Driver变量
func driver(c car) {
	c.run()
}
func (f falali) run() {
	fmt.Println(f.brand, "\t速度1000")
}

func (d dazhong) run() {
	fmt.Println(d.brand, "\t速度20000")
}

func main() {
	f1 := falali{
		brand: "法拉利",
	}
	driver(f1) //同时调用 driver方法
	da := dazhong{
		brand: "大众",
	}
	driver(da)

}
