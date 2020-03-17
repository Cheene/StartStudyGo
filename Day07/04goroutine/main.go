package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

//程序启动之后会创建一个主 goroutine 去执行
func main() {
	for i := 0; i < 1000; i++ {
		//go  hello(i) //开启一个单独的 goroutine 执行 hello 函数
		go func(i int) {
			fmt.Println("main :", i)
		}(i) //保证仅有一次
	}
	time.Sleep(1 * time.Second)
	//main函数结束了 释放了所有的资源
	fmt.Println("main")

}
