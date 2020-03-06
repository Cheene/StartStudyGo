package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

/**
延迟加载，注册的时候就已经进去了
*/
//1. a:=1
//2. b:=2
//3. defer calc("1",1,calc("10",1,2)
//4. calc("10",1,2)
