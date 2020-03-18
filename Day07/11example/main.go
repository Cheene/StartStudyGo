package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**

 */
type Job struct {
	x int64
}

type Result struct {
	job *Job
	sum int64
}

var jobChan = make(chan *Job, 100)
var resultChan = make(chan *Result, 100)

func product(jobs chan<- *Job) {
	defer wg.Done()
	for {
		temp := rand.Int63()
		ret := &Job{x: temp}
		jobs <- ret
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(jobs <-chan *Job, resultes chan<- *Result) {
	defer wg.Done()
	for {
		job, ok := <-jobs
		if !ok {
			break
		}
		sum := int64(0)
		n := job.x
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		newResult := &Result{
			job: job,
			sum: sum,
		}
		resultes <- newResult
	}
}

var wg sync.WaitGroup

func main() {

	wg.Add(1)
	go product(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go consumer(jobChan, resultChan)
	}

	//输出
	for result := range resultChan {
		fmt.Printf("value:%d,sun:%d\n", result.job, result.sum)
	}
	wg.Wait()

}
