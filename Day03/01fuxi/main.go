package main

import (
	"fmt"
)

func main() {
	var name string
	name = "Chenene"
	fmt.Print(name)

	var age [30]int
	age = [30]int{1, 2, 3, 4}
	fmt.Println(age)

	var ag2 = [...]int{12, 3, 12412412, 4, 21}
	fmt.Println(ag2)

	var ag3 = [...]int{1: 100, 12: 3000}
	fmt.Println(ag3)

	//二维数组 最外层才可以使用...
	var a1 = [...][3]int{
		[3]int{1, 2, 3},
		[3]int{3, 4, 6},
	}
	fmt.Println(a1)
	//数组时值类型
	x := [3]int{0, 1, 3}
	fmt.Println(x)
	//向函数中传递的是一个副本
	f1(x) // 仍然是123 ，Go语言中的幻术传递的都是值

	fmt.Println(x)

	y := x         // y 初始化的时候仍然是 x 的副本，
	y[1] = 100     //修改 y 的数值对 x 并不会由任何的影响
	fmt.Println(x) // y 是 X 的副本，x 不会变
	// 切片是引用类型，底层是数组存值，切片本身并不存储数值。
	// 切片的底层是数组，只有数组存储值，而切片存储
	// 关于切片 slice: 仅能装相同的元素
	var s1 []int
	fmt.Println(s1)        // [] 本质上是 nil ,并没有分配内存，即
	fmt.Println(s1 == nil) //
	s1 = []int{1, 2, 3}    // 切片初始化01
	fmt.Println(s1)
	// 初始化，已经分配了内存
	s2 := make([]bool, 2, 4)
	fmt.Println(s2 == nil)
	fmt.Println(s2)
	// false;
	s3 := make([]int, 0, 4)               // 4 是底层数组的长度。0是切片本身的长度。长度小于容量；指针是指向数组的首位置
	fmt.Println("s3 == nil? ", s3 == nil) // 初始化后一定有一个指针指向内存地址，底层数组的容量是4，切片本身的长度为0 而已
	// 切片的赋值与拷贝

	ss1 := []int{1, 2, 3} // [1,2,3]
	ss2 := ss1            // [1,2,3]
	fmt.Println(ss2)
	// 切片不存储值，类似于指针指向同一个内存首地址
	ss2[1] = 200
	fmt.Println(ss1) // [1,200,3]
	fmt.Println(ss2) // [1,200,3]

	//var slice1 []int // 这里没有分配内存
	//slice1[0] = 100 //index out of range [0] with length 0
	//fmt.Println(slice1)

	var slice1 []int
	slice1 = append(slice1, 233)
	fmt.Println(slice1)

	//切片的Cope
	//var s4 []int // nil []
	//var s4 = make([]int,0,3)//长度是0 []
	var s4 = make([]int, 3, 3) //这里是值传递
	s5 := s1
	copy(s4, s1) // copy 是数值的 copy，属于底层数组的副本
	fmt.Println(s2)
	s5[1] = 1000
	fmt.Println(s1)
	fmt.Println(s5)
	fmt.Println(s4)

	//Go 指针只能读不能修改，不能修改指针变量指向的内存地址
	name = "Chenene" // 不能使用 := ,No new variables on left side of :=
	nameP := &name
	fmt.Println(nameP)
	fmt.Printf("%T\n", nameP)
	nameValue := *nameP
	fmt.Println(nameValue)
	// Map 也是需要分配内存的
	var m1 map[string]int
	fmt.Println(m1 == nil) // true,所以需要初始化
	m1 = make(map[string]int, 10)
	m1["Chenene"] = 100
	fmt.Println(m1)
	fmt.Println(m1["Chen"]) // IF KEY is not exist，不会报错，返回值与value相关
	score, ok := m1["Chenene"]
	if !ok {
		fmt.Println("Not Exist")
	} else {
		fmt.Printf("Exist : %d\n", score)
	}

	delete(m1, "Chen") // 删除的 KEY不存在，什么也不干
	fmt.Println(m1)
	delete(m1, "Chenene")
	fmt.Println(m1)
	fmt.Println(m1 == nil)
}

func f1(a [3]int) {
	a[1] = 100
}
