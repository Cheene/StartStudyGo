package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = sync.Map{}
var wg sync.WaitGroup

func main() {
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			m.Store(key, i)
			value, ok := m.Load(key)
			if !ok {
				fmt.Printf("This is no value,key: #%v\n", value)
			}
			fmt.Printf("m[%v] = '%v'\n", key, value)
			wg.Done()
		}(i)
		wg.Wait()
	}
}
