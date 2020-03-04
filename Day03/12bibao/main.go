package main

import "fmt"

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := calc(10)
	//每次打印的时候，修改的都是同一个base
	fmt.Println(f1(1), f2(3)) // 11，8
	fmt.Println(f1(5), f2(7)) // 13，6
	fmt.Println(f1(4), f2(2)) //  10，8
}
