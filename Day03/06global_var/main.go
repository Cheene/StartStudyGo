package main

import "fmt"

//变量的作用域
var x = 100

//定义函数
// 函数中变量的查找顺序
// 1.先在函数内部进行查找；
// 2.找不到就向函数外部查找，一直找到全局
func f1() {
	fmt.Println(x)
}

func f2() {
	x := 2
	fmt.Println(x)
}

func main() {
	f1()
	f2()

	//语句块作用域
	if i := 10; i < 18 {
		fmt.Println("乖乖上学")
	} else {
		fmt.Println("乖乖不上学")
	}
	//fmt.Println(i) //Unresolved reference 'i'

}
