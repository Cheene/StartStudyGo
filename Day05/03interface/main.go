package main

import "fmt"

//引出接口的实例
type speaker interface {
	speak() // 只要实现了 speak 方法的变量都是speaker类型 ===》方法签名
}

type cat struct {
}

type dog struct {
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}
func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func da(x speaker) {
	x.speak()
}

func main() {
	var c1 cat
	var d1 dog
	da(c1)
	da(d1)

	var ss speaker
	ss = c1 // c1 实现了ss的所有的方法
	fmt.Println(ss)

}
