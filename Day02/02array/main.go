package main

import "fmt"

// 数组 存放元素的容器{类型和容量}

func main() {
	//	数组的长度是数组类型的一部分
	var a1 [2]bool
	a1[0] = true
	a1[1] = false

	for _, b := range a1 {
		fmt.Println(b)
	}
	fmt.Printf("%T", a1)
	var a2 [3]bool
	//初始化1
	fmt.Println(a2)
	a2 = [3]bool{true, true, true}
	//初始化2
	// a100 := [100]{1,2,3,4,5,6,7,89,9}
	// 根据初始值自动推断数组的长度是多少
	a10 := [...]int{1, 23, 12, 312, 4, 12, 4, 12}
	//初始化方式3
	a11 := [5]int{1, 2}
	fmt.Println(a10)
	fmt.Println(a11)
	//根据素银初始化
	a12 := [5]int{0: 1, 4: 22}
	fmt.Println(a12)

	citys := [...]string{"北京", "上海", "杭州"}

	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}

	a111 := [3][2]int{
		[2]int{1, 2},
		[2]int{5, 6},
		[2]int{7, 4},
	}

	for i, v := range a111 {
		fmt.Println(i, v)
	}
	// 数组是非指针类型，属于值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b1, b2) // [1 2 3] [100 2 3] 拷贝的，拷贝之后改完成，之前的不会修改

	//练习一
	sum := 0
	tempSum := [...]int{1, 3, 5, 7, 8}

	for _, val := range tempSum {
		sum += val
	}
	fmt.Printf("Sum : %d\n", sum)
	//练习二
	for i := 0; i < len(tempSum); i++ {
		for j := i + 1; j < len(tempSum); j++ {
			if tempSum[i]+tempSum[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
			if tempSum[i]+tempSum[j] > 8 {
				break
			}
		}
	}
}
