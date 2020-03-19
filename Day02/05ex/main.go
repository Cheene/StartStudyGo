package main

import (
	"fmt"
	"strings"
)

//统计单词的个数 map函数，哈希函数算法
/*
	1 将 str 以空格为分割，转换为切片;
	2 定义一个 map
	3 循环遍历切片并且插入到 map
	4 遍历 map
*/
func main() {
	str := "how do you do"

	splitStr := strings.Split(str, " ")
	fmt.Println(splitStr)
	mapSlic := make(map[string]int, len(splitStr))

	for _, value := range splitStr {
		t, ok := mapSlic[value]

		if !ok {
			t = 0
		}
		t += 1
		mapSlic[value] = t
	}

	//
	for key, v := range mapSlic {
		fmt.Println(key, ":", v)
	}
}
