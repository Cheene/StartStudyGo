package main

import (
	"fmt"
	"sync"
)

//需要指定通道中元素的类型，是一个引用类型，所以需要开辟空间
var a chan int
var wg sync.WaitGroup

func bufferChannel() {
	a = make(chan int, 16) // 带缓冲区通道
	a <- 10
	fmt.Println("BufferChannel:10发送到通道中了")
}

func noBufferChannel() {
	//通道必须使用make函数初始化才能够使用
	a = make(chan int) // 不带缓冲区通道
	wg.Add(1)
	//a <- 10 // 卡住了，没有接的
	//为什么没有接的？？
	go func() {
		defer wg.Done()
		//只有把值放进去的时候，才把值取出来
		x := <-a
		fmt.Printf("NOBufferChannel::后台 goroutine 获取%d 成功了\n", x)
	}()
	a <- 100
	fmt.Println(a)
	wg.Wait()

}

func main() {
	noBufferChannel()
	bufferChannel()
	close(a)
}
