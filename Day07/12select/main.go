package main

import "fmt"

// select 随机的执行case 语句
// 没有 {}语句的 select 会一直等待
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
