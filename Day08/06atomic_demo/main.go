package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//原子操作
var x int64
var lock sync.Mutex

func add() {
	//lock.Lock()
	//x++ // 两步操作，第一步获取；第二步 +1
	//lock.Unlock()
	atomic.AddInt64(&x, 1)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Printf("%d\n", x)
}
