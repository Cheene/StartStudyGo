package main

import (
	"fmt"
	"sync"
	"time"
)

//PV读写锁
/*
读的时候没有写的时候才可以写
写的时候直接写
*/
var (
	x      = 0
	lock   sync.Mutex
	rwLock sync.RWMutex
	wg     sync.WaitGroup
)

func read() {
	defer wg.Done()
	rwLock.RLock() //这里是读锁
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
}

func main() {
	now := time.Now()
	//先写
	for i := 0; i < 100; i++ {
		go write()
		wg.Add(1)
	}
	//再读
	for j := 0; j < 1000; j++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	//输出时间
	fmt.Println("时间： ", time.Now().Sub(now))
}
