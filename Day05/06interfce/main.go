package main

import "fmt"

//接口的实现

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
	fmt.Println("小鸡吃虫子")
}

func (c cat) move() {
	fmt.Println("走猫步，乱跳跳")
}

//实现的不是一个同方法，不满足接口的定义
//func (c cat)eat(){ // 不带参数的，不是一个方法
//	fmt.Println("猫吃鱼")
//}

func (c cat) eat(some string) {
	fmt.Println(some)
}

func main() {
	var a1 animal
	var c1 cat
	c1 = cat{
		name: "蓝猫",
		feet: 4,
	}
	a1 = c1
	c1.move()
	fmt.Println(a1)

	kfc := chicken{
		name: "KFC",
		feet: 2,
	}

	a1 = kfc
	fmt.Printf("%T\n", a1)
	a1.eat("小")
	a1.move()
}
