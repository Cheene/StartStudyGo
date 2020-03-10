package main

import "fmt"

//空接口作为参数
func show(a interface{}) {
	fmt.Printf("type: %T\tValue:%v\n", a, a)
}

func main() {
	// interface 是关键字
	// interface{} ：空接口类型
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)

	m1["name"] = "Chene"
	m1["age"] = 1000
	m1["hobby"] = [...]string{"唱跳", "跑"}

	fmt.Println(m1) //无序

	show(false)
	show("Chenene")
	show(nil)
	show(m1)
	//show(interface{})
}
