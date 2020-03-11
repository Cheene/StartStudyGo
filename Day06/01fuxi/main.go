package main

import "fmt"

//类型断言

func main() {
	var a interface{}
	a = 45
	fmt.Printf("%T\n", a)

	if v, ok := a.(int8); !ok {
		fmt.Printf("猜对了%v", v)
	} else {
		fmt.Printf("猜错了%v", v)
	}

}
