package main

import (
	"fmt"
	"strconv"
	"sync"
)

//Go 内置的map 不是并发安全的
// concurrent map writesK = 5,value = 5
// ? 为什么会这样？？
// 所以在写操作的时候，加锁
var m = make(map[string]int)

var lock sync.RWMutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			lock.Lock()
			set(key, n)
			lock.Unlock()
			fmt.Printf("K = %v,value = %v\n", key, get(key))
			wg.Done()
		}(i)
	}

	wg.Wait()
}
