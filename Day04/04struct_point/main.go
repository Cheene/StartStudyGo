package main

import "fmt"

type person struct {
	name  string
	age   int
	hobby []string
}

//go 语言中函数传参永远是拷贝
func f(x person) {
	x.age = 10 // 修改的是副本的gender
}

func f2(x *person) {
	//(*x).age = 10
	x.age = 10 // 语法糖，自动根据指针寻找对应的变量
}

func main() {
	var p person
	p.age = 100
	fmt.Println(p)
	f(p)
	fmt.Println(p)
	f2(&p)
	fmt.Println(p)

	//结构体指针1
	var p2 = new(person)
	p2.name = "lixiang"
	fmt.Printf("%T : %v : %p\n", p2, p2, p2)
	//结构体指针2
	var p3 = person{"chenene", 19, []string{"篮球", "乒乓球"}}
	fmt.Println(p3)
	fmt.Printf("%T\n", p3)
	// key - value 方式初始化
	p4 := person{
		name:  "aaa",
		age:   1,
		hobby: []string{},
	}
	fmt.Printf("%T\n", p4)
	fmt.Printf("%v\n", p4)

}
