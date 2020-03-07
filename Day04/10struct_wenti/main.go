package main

import "fmt"

//1 myInt(100) 是什么东西 不是函数
type myInt int

//不能给别的包类型变量添加方法，只能给自己的包的类型添加方法
func (m myInt) hello() {
	fmt.Println("Hello int")
}

//问题2 结构体的初始化
type person struct {
	name string
	age  int
}

func main() {
	//问题一
	var x int32 = 10
	fmt.Println(x)

	var y int32
	y = 100
	fmt.Println(y)

	var x1 = int(222)
	fmt.Println(x1)

	// 声明一个myInt类型的变量 m ，初始化数值为 100
	var m myInt
	m = 10
	fmt.Println(m)

	var m1 myInt = 12
	fmt.Println(m1)

	m22 := myInt(1000) //强制类型转换
	fmt.Println(m22)

	//
	s1 := []int{1, 2}
	map1 := map[string]int{
		"stu1": 100,
		"stu2": 200,
	}
	fmt.Println(s1, map1)
	// 法1
	var p person
	p.name = "Chenene"
	p.age = 10
	fmt.Println(p)
	// key - value 类型； 不需要 new
	p1 := person{
		name: "Chen",
		age:  158,
	}
	fmt.Println(p1)

	//值列表初始化
	var p3 = person{
		name: "JJJ",
		age:  0,
	}
	fmt.Println(p3)

}

//为什么有构造函数 -- 别人调用我，我能给她一个Person类型的变量；
//必须先有一个Person
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

func newPerson2(name string, age int) *person {
	//*person 类型是指针的， &person 是返回person的内容，取地址的内容
	return &person{
		name: name,
		age:  age,
	}
}
