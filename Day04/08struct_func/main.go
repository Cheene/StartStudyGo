package main

import "fmt"

type dog struct {
	name string
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法是作用于特定类型的函数
//接收者表示调用方法具体的类型，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s 汪汪汪\n", d.name)
}

func main() {
	d1 := newDog("LISHI")
	d1.wang()
}
