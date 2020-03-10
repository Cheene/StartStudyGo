package main

import "fmt"

//使用指针解咒着实现了接口的办法
type animal interface {
	move()
	eat(s string)
}

type cat struct {
	name string
	feet int8
}

type chicken struct {
	name string
	feet int8
}

func (c chicken) move() {
	fmt.Println("小鸡快跑")
}

func (c chicken) eat(some string) {
	fmt.Println(some, "小鸡吃虫子")
}

//使用指针解咒着实现了接口的办法
func (c *cat) move() {
	fmt.Println("走猫步，乱跳跳")
}

//实现的不是一个同方法，不满足接口的定义
//func (c cat)eat(){ // 不带参数的，不是一个方法
//	fmt.Println("猫吃鱼")
//}

func (c *cat) eat(some string) {
	fmt.Println(some)
}

func main() {

	var a1 animal
	c1 := cat{"tom", 4}
	c2 := &cat{"假老练", 4}

	a1 = &c1 //指针接收者
	fmt.Println(a1)
	a1 = c2
	fmt.Println(a1)
}
