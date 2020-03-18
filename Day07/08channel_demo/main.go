package main

import (
	"fmt"
	"sync"
)

// channel 练习
// 1 启动一个 goroutine，生成100个数发送到ch1
// 2 启动一个 goroutine，从ch1张取值并计算平方值放到 ch2中

var wg sync.WaitGroup
var once sync.Once

// chan<- 代表着只能向通道中发送值，
func f1(ch1 chan<- int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

//ch1仅能取值，ch2仅能放值
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()

	//for x := range ch1{ 这样或产生错误，因为 ch1 还不一定能够生成成功。
	//
	//}

	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	//确保仅仅执行一次
	once.Do(func() { close(ch2) })
}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)
	wg.Add(3)
	go f1(ch1)
	go f2(ch1, ch2)
	go f2(ch1, ch2)

	for ret := range ch2 {
		fmt.Println(ret)
	}
	wg.Wait()

}
