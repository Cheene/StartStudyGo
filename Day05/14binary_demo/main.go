package main

import "fmt"

//二进制的位数
const (
	EAT   int = 4
	SLEEP int = 2
	DA    int = 1
)

//二进制实用途
///111 1- 吃饭 1- 睡觉 1- 打豆豆
func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(EAT | DA)         // 101
	f(EAT | DA | SLEEP) // 101
}
