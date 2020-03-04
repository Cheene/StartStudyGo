package main

import (
	"fmt"
	"unicode"
)

/**
统计字符中汉字的数量
*/
func main() {
	s1 := "你是Chenene吗？"
	//记录个数
	var count int
	// 1 遍历字符串
	for _, c := range s1 {
		// 2 判断是否是汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	// 3 输出汉字的个数
	fmt.Println(count)

	//回文数的判断，字符串的逆转与字符串相等
	s2 := "上海自来水来自海上"
	//字符串转换成切片
	r := make([]rune, 0, len(s2))
	for _, c := range s2 {
		r = append(r, c)
	}
	flag := true
	for i, _ := range r {
		if i > len(r)/2 {
			break
		}
		if r[i] != r[len(r)-i-1] {
			flag = false
			break
		}
	}
	fmt.Println(r) // 中文占三个字节
	fmt.Println(flag)

}
