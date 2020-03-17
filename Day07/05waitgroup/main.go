package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//waitGroup

func f() {
	//rand.Seed()
	for i := 0; i < 100; i++ {
		rand.Seed(time.Now().UnixNano()) //设置随机数种子，保证编译后每一次的执行都不一样
		r1 := rand.Int63n(3000)          // 0 <= x < 3000
		r2 := rand.Int()

		fmt.Println(r1, ":", r2)
	}
}

func f1(i int) {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(10000)))
	fmt.Println(i)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go f1(i)
	}
	//如何知道这10个goroutine 都结束？
	wg.Wait() //等待 wg 计数器 减为 0
}
