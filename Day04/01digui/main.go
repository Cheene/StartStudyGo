package main

import "fmt"

//递归
// 递归函数和退出条件。适合问题相同，问题规模越来越小的问题
//永远不要高估自己

//阶乘
func jiecheng(num int) int {
	if num <= 1 {
		return 1
	}
	return num * jiecheng(num-1)
}

// 上台阶的问题：一次可以走一步，一次走两步。有多少种走法
func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-2) + taijie(n-1)
}
func main() {
	fmt.Println(jiecheng(6))
	fmt.Println(taijie(4))
}
