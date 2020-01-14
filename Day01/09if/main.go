package main

import "fmt"

func main() {
	//if的特殊使用。 作用域
	if age := 19; age < 18 {
		fmt.Println("18")
	} else {
		fmt.Println("28")

	}
	//for 循环
	for i := 20; i < 100; i += 10 {
		fmt.Println(i)
	}
	//无限循环
	// for {
	// 	fmt.Println("123")
	// }

	// for range 主要用于
	a := "[]rune()"
	for index, value := range a {
		fmt.Printf("%v : %v \n", index, value)
	}
}
