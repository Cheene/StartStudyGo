package main

//接口可以嵌套
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(food string)
}

//同一个结构体可以实现多个接口
type cat struct {
	name string
	feet int8
}

func (c *cat) move() {

}

func (c *cat) eat() {

}

func main() {

}
