package main

import (
	"fmt"
	"sync"
)

//测试是否是因为make的原因，结果不是。
var wg sync.WaitGroup

//var lock sync.Mutex
func main() {
	slice := make([]int, 400)

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			//lock.Lock()
			slice[i] = i
			//lock.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for index, value := range slice {
		fmt.Printf("[%d] : %d\n", index, value)
	}
}
