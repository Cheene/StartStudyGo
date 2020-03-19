package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("%d \n", a)
	fmt.Printf("%b \n", a) // b binary 二进制

	//
	var b int = 077
	fmt.Printf("%o \n", b) // o 八进制
	fmt.Printf("%d \n", b)
	//
	fmt.Printf("%x \n", b)
	fmt.Printf("%T \n", b)
	//强制类型
	c := int8(9)
	fmt.Printf(" c : %T\n", c)
}
