package main

import "fmt"

//匿名字段 适合于字段少比较简单的场景
type person struct {
	string
	int
}

func main() {
	p1 := person{
		string: "指针",
		int:    12,
	}

	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
