package main

import (
	"fmt"
	"strconv"
)

//strconv

func main() {

	//int 转 string
	i := (21306)
	ret2 := fmt.Sprintf("%d", i)
	fmt.Printf("%#v\n", ret2)
	i2a := strconv.Itoa(i)
	fmt.Println(i2a)

	//字符串中解析成int类型
	str := "1000"
	istr, _ := strconv.Atoi(str)
	fmt.Printf("%#v\n", istr)

	ii, err := strconv.ParseInt("1263152", 10, 64)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Printf("%#v\n", ii)
	//字符串中解析出布尔值

	boolStr := "true"
	booValue, _ := strconv.ParseBool(boolStr)
	fmt.Println(booValue)

	//浮点数
	floatStr := "1.34"

	floatValue, _ := strconv.ParseFloat(floatStr, 64)
	fmt.Println(floatValue)
}
