package main

import (
	"fmt"
)

//关闭通道
func main() {
	ch1 := make(chan int, 2)
	ch1 <- 10
	ch1 <- 20
	close(ch1)

	for x := range ch1 {
		fmt.Println(x)
	}
	//对已经关闭的通道进行重新读取是可以的。
	x, ok := <-ch1
	fmt.Println(x, ok)

}
