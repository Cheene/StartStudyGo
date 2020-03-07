package main

import "fmt"

//构造函数
type person struct {
	name string
	age  int
}

type dog struct {
	name string
}

//值就是拷贝，值接收者
func (p person) guonian() {
	p.age++
}

//地址接收者，传内存地址进去
// 1 修改接收者的值； 2 拷贝代价比较大的时候；3 保证一致性的时候
func (p *person) zhenguonian() {
	p.age++
}

// 保证一致性
func (p *person) dream() {
	fmt.Println("一直增加睡后收入")
}

//返回的是结构体还是结构体指针 如果体量大，直接传递指针类型
//当结构体比较大的时候尽量使用结构体指针，减少程序的开销
// 约定： 用 new 开头表示构造函数
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//当结构体比较大的时候尽量使用结构体指针，减少程序的开销
func newPerson2(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("元帅", 18)
	p2 := newPerson("Chenene", 20)
	fmt.Println(p1, p2)
	p3 := newDog("XXX")
	fmt.Println(p3)
	fmt.Printf("%d \t", p1.age)
	p1.guonian()
	fmt.Printf("%d \t", p1.age)
	p1.zhenguonian()
	fmt.Println(p1.age)
	p1.dream()
}
