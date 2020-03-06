package main

import "fmt"

func main() {
	fmt.Print("Chenene")
	fmt.Print("CCC")
	/*
		Printf
		%T 查看类型
		%p 指针
		%v 原样输出值
		%f 浮点
		%t 布尔值
		%U Unicode格式
		%q
	*/

	m1 := make(map[string]int, 1)

	m1["HEE"] = 2220

	fmt.Printf("%v\n", m1)
	fmt.Printf("%+v\n", m1)
	fmt.Printf("%#v\n", m1)
	num := 90
	fmt.Printf("%d%%\n", num)
	// 整数转字符
	fmt.Printf("%q", 65)
	//浮点数
	fmt.Printf("%b\n", 3.312412512)
	//字符串加双引号
	fmt.Printf("%q\n", "lxxx")

	n := 12.3455765
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.5f\n", n)
	fmt.Printf("%4.3f\n", n)
	fmt.Printf("%9.f\n", n)

	//	获取用户输入
	//	var s string
	//fmt.Scan(&s)
	//fmt.Println(s)

	//var(
	//	name string
	//	age int
	//	class string
	//)
	//
	//fmt.Scanf("%s %d %s\n",&name,&age,&class)
	//fmt.Printf("%s %d %s\n",name,age,class)
}
