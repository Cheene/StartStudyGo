package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Printf("wroker:%d start job:%d\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, i)
		results <- i * 2
		notifyCh <- struct{}{} // struct{}{} 匿名结构体的实例化
	}
}

var wg sync.WaitGroup
var notifyCh = make(chan struct{}, 5)

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启五个任务,即向 jobs中输入五个值
	go func() {
		for v := 1; v <= 5; v++ {
			jobs <- v
		}
		close(jobs)
	}()
	//开启三个 goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	//判断results是否完成
	go func() {
		for i := 0; i < 5; i++ {
			<-notifyCh
		}
		close(results)
	}()
	//此时 results 已经满了，开始输出
	for x := range results {
		fmt.Println(x)
	}

	//走到这里仅仅是全部启动 worker,并不代表全部执行完成

}
