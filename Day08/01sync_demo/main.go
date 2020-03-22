package main

import (
	"fmt"
	"sync"
)

//锁 sync
//公共的存在的问题
var x = 0

//解决方式 -- 锁，加在对公共资源操作的地方
var lock sync.Mutex

var wg sync.WaitGroup

func add() {
	for i := 0; i < 50000; i++ {
		//这里有三步：
		//1 取；2 +1 ；3 重新赋值
		//x = x + 1
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}

	wg.Done()

}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
