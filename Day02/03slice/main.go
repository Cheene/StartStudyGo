package main

import "fmt"

//切片 类型相同，长度可变 底层为数组，上层为封装

func main() {
	//切片的定义
	var slices []int
	var stringss []string

	fmt.Println(slices)
	fmt.Println(cap(slices))
	fmt.Println(stringss)
	fmt.Println(stringss == nil) //没有初始化就是没有分配内存
	fmt.Println(slices == nil)

	stringss = []string{"沙河", "张江", "平山村"}
	fmt.Println(stringss)
	// 切片的容量 	//cap函数返回的是分配给切片的动态内存大小
	fmt.Println(cap(stringss))
	//2 由数组通过切割得到切片
	// 从你切片的第一个去算  算到底层数组的最后一个就是他的容量
	a1 := [...]int{1, 2, 312, 321, 3, 124, 12, 3, 21}

	s22 := a1[0:2] // 左闭右开

	s23 := a1[:3] // 1 2 312, 仅有三个元素
	s24 := a1[4:] //3 124 ...
	s25 := a1[:]  // all
	fmt.Println(a1[0])
	fmt.Println("a1: ", a1)
	fmt.Println("s22: ", s22)
	fmt.Println("s23: ", s23)
	fmt.Println("s24: ", s24)
	fmt.Println("s25: ", s25)
	//切片的容量是指底层数组的容量,是从切片的第一个元素到最后一个元素
	//他们两个的区别是什么？
	//在什么样的情况下会，由数组获得的切片会修改大小
	fmt.Printf("len(s22):%d,cap(s22):%d\n", len(s22), cap(s22))
	fmt.Printf("len(s24):%d,cap(s24):%d\n", len(s24), cap(s24))
	//切割再切割
	s26 := s24[3:]
	fmt.Printf("len(s26):%d,cap(s26):%d\n", len(s26), cap(s26))
	//切片作为引用类型，同时指向同一个底层数组。
	s26[0] = 10000
	fmt.Println(a1)
	fmt.Println(s26)

	//make函数与切片

}
