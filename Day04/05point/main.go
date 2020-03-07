package main

import "fmt"

func main() {
	var a int
	a = 100
	b := &a

	fmt.Printf("a: %T, b:%T\n", a, b)
	fmt.Printf("a: %v, b:%v\n", &a, b)

}
