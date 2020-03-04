package main

import "fmt"

//defer 栈
/**
比如 socket ， 数据库连接等用来释放资源的地方。先通过 defer 释放资源

*/

func deferDemo() {
	fmt.Println("start")
	fmt.Println("ing")
	fmt.Println("end")
	fmt.Println()

}
func deferDemo2() {
	fmt.Println("start")
	defer fmt.Println("ing") // 把后面的语句延迟到函数即将返回的时候载执行
	fmt.Println("end")
	fmt.Println()
}
func deferDemo3() {
	fmt.Println("start")
	defer fmt.Println("ing") // 把后面的语句延迟到函数即将返回的时候载执行
	fmt.Println("end")
	defer fmt.Println()
}
func main() {
	deferDemo()
	deferDemo2()
}
