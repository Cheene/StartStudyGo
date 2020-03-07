package main

import "fmt"

//结构体模拟实现其他语言的 继承

type animal struct {
	name string
}

func (a animal) move() {
	fmt.Printf("%s 会动\n", a.name)
}

//狗
type dog struct {
	feet   uint8
	animal // animal 拥有的方法，dog此时也有了
}

//汪汪汪
func (d dog) wang() {
	fmt.Printf("%s 汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		feet: 4,
		animal: animal{
			"zz",
		},
	}
	d1.wang()
	d1.move()
}
