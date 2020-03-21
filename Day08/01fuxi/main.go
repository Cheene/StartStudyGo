package main

import (
	"fmt"
	"math/rand"
	"time"
)

//channel

func sendNum(ch chan<- int) {
	for {
		i := rand.Int()
		ch <- i
		time.Sleep(5 * time.Second)
	}
}

func main() {
	ch := make(chan int, 1)
	ch <- 100
	<-ch //取出来了
	//阻塞了，在等着取值
	//x,ok := <-ch //这里没有值 fatal error: all goroutines are asleep - deadlock!
	//fmt.Print(x,ok)
	go sendNum(ch)
	for {
		x, ok := <-ch
		fmt.Println(x, ok)
	}
}
