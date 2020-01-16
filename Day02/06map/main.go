package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%2d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}
	// 取出 map 中所有的key,存入到切片keys。
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//排序
	sort.Strings(keys)
	//
	fmt.Println("按照学生学号")
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
	//按照成绩排序
	var values = make([]int, 0, 200)
	for _, value := range scoreMap {
		values = append(values, value)
	}
	sort.Ints(values)
	fmt.Println("按照学生成绩")
	for _, key := range values {
		for k, v := range scoreMap {
			if v == key {
				fmt.Println(k, ":", v)
			}
		}
	}
	// map 和 slice 的组合
	// 一个 map类型的切片
	var ms1 = make([]map[int]string, 3, 10) // 切片的长度不能为0

	// ms1[0][1] = "A" // index out of range [0] with length 0

	// 1 先初始化map
	ms1[0] = make(map[int]string, 1)
	ms1[0][10] = "chenene"
	fmt.Println(ms1)

	//一个 切片类型的 map
	var ms2 = make(map[string][]int, 10) // 一个 map key: string value : map类型

	ms2["00"] = make([]int, 0, 10)

	value2, ok := ms2["00"]
	if !ok {
		value2 = make([]int, 0, 2)
	}
	value2 = append(value2, 1, 2)
	ms2["00"] = value2

	fmt.Println(ms2)
}
