package main

import "fmt"

func main() {
	var m1 map[string]int // key: string , value:int
	//assignment to entry in nil map 还未有初始化
	m1 = make(map[string]int, 10)
	m1["山东"] = 100

	fmt.Println(m1)
	v, ok := m1["这款"]
	if !ok {
		fmt.Println("没有此key值")
	} else {
		fmt.Println(v)
	}

	m1["xxx"] = 12

	// for range 打印键值对
	for k, v := range m1 {
		fmt.Println(k, ":", v)
	}
	//删除
	delete(m1, "xxx")
	fmt.Println(m1)
}
