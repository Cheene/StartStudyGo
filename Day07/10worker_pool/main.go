package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Printf("wroker:%d start job:%d\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job:%d\n", id, i)
		results <- i * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启三个 goroutine
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	//开启五个任务,即向 jobs中输入五个值
	for v := 1; v <= 50; v++ {
		jobs <- v
	}
	close(jobs)
	//输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
