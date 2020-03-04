package main

import "fmt"

//func f1(参数) 返回值类型{
//
//}

func f1() {
	fmt.Println("NIHAO Chenene")
}
func f2(name string) {
	fmt.Printf("Name : %s\n", name)
}
func f3(x int, y int) int {
	return x + y
}

//擦描述类型的简写
func f4(x, y int) int {
	return x + y
}

//可变参数
// y 是int 类型的切片
func f5(title string, y ...int) []int {
	return y
}

//返回值命名
func f6(x, y int) (sum int) {
	sum = x + y // 使用命名的函数值，可以直接使用返回值的变量， return后面可以直接省略返回值的变量
	return      // 即返回的就是 sum
}

//多个返回值
func f7(x, y int) (sum int, sub int) { // 要么全写，要么全不写
	sum = x + y
	sub = x - y
	return
}

func main() {
	f1()
	f2("Chenene")
	f2("Nuoning")
	sum := f3(1, 3)
	fmt.Println(sum)
	fmt.Println(f4(100, 300))

	index := f5("Chenen", 213, 213, 12, 312, 32, 14234, 32, 523, 432)
	fmt.Println(index)
	fmt.Println(f6(1, 2))

	sum, sub := f7(3, 4)
	fmt.Println(sub)
	fmt.Println(sum)
}
