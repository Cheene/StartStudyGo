package main

import "fmt"

// painc 和 recover 成对出现

func funcA() {
	fmt.Println("A")
}

func funcB() {
	//可能问题： 刚刚打开了数据库的连接，需要释放资源
	defer func() {
		//
		fmt.Println("释放数据库连接")
		//尝试恢复
		err := recover()
		fmt.Println(err)
	}()
	//释放数据库连接
	//出现严重的错误
	panic("出现严重的错误")
	//这里以后不会执行
	fmt.Println("B")
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
