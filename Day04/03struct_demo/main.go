package main

import "fmt"

//结构体
// 多种维度的数据组合在一起
// type 后面就是类型
type person struct {
	name  string
	age   int
	hobby []string
}

func main() {
	var chenene person
	chenene.name = "Chenene"
	chenene.age = 24
	chenene.hobby = []string{"basketball", "sing"}
	fmt.Println(chenene)

	//访问变量P的字段
	fmt.Println(chenene.name)
	fmt.Printf("tyep: %T value:%v\n", chenene, chenene)
}
