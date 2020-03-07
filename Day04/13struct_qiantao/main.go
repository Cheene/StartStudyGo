package main

import "fmt"

//嵌套，结构体

type address struct {
	province string
	city     string
}

type workPlace struct {
	province string
	city     string
}
type person struct {
	name string
	age  int
	addr address
}

type company struct {
	p       person
	address // 匿名嵌套结构体
	workPlace
}

func main() {
	p1 := person{
		name: "周琳",
		age:  12,
		addr: address{
			"山东",
			"威海",
		},
	}

	fmt.Println(p1.name)
	fmt.Println(p1.addr.province)
	//fmt.Println(p1.city)// 失败
	c1 := company{
		p: p1,
		address: address{
			"山东",
			"莱芜",
		},
	}
	// 1 现在自己的结构体中找；2 再去匿名嵌套的结构体使用
	//fmt.Println(c1.city) //Error 匿名结构体存在错误
	fmt.Println(c1.address.city) // 字段冲突解决
}
