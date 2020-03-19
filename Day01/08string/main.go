package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Hello,诺宁"
	c1 := 'A'
	fmt.Printf("%d \n", c1)
	fmt.Printf("%s \n", s)
	// \r 回车符

	path := "\"D:\\Go\\src\\github.com\\Chenene\\StartStudyGo\\Day01\""
	fmt.Printf("path : %s \n", path)
	ss := "I'm OK"
	fmt.Printf("path : %s\n", ss)
	//多行的字符串
	s2 := `
		诗情额
		雨落黄昏花易落
	`
	fmt.Printf("%s\n", s2)
	fmt.Printf("len("+s2+") : %d\n", len(s2))
	//拼接
	name := "china"
	word := "I lova"
	fmt.Println(name + " " + word)
	// Sprintf,保存格式化的参数到变量中
	ss1 := fmt.Sprintf("%s %s", s, word)
	fmt.Println(ss1)
	//字符串的分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(s2, "诗"))
	//测试前缀或者后缀以某个字符串结束
	// 前缀
	fmt.Println(strings.HasPrefix(s2, "古诗"))
	// 后缀
	fmt.Println(strings.HasSuffix(s2, "花易落"))
	//
	s4 := "abcdef"
	fmt.Println("A: ", strings.Index(s4, "b"))
	fmt.Println(strings.LastIndex(s4, "f"))

	for _, c := range s2 {
		fmt.Printf("%c \n", c)
	}

	// rune 类型 中文，日文等非英文推荐的类型 是 int32 的别名< 一个中文占3B=24bit >
	//string 是不能修改的变量
	s22 := "白萝卜"
	// s22[0] = "红" Error
	s3 := []rune(s22) // 强制转换成切片的类型
	s3[0] = '红'
	fmt.Println(string(s3)) // 红萝卜
	// 字符串与字符的区别
	c3 := 'H'
	c4 := "H"
	c5 := byte('H') // byte <=> unit8
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("c5:%T\n", c5)
	//类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Printf("%v \n", f)

	//
	var a int32
	var b float64
	var c bool
	var d string

	fmt.Printf("a ： %T : %v \n", a, a)
	fmt.Printf("b ： %T : %v \n", b, b)
	fmt.Printf("c ： %T : %v \n", c, c)
	fmt.Printf("d ： %T : %v \n", d, d)

	str := "hellozZ沙河小王子"
	rus := []rune(str)

	count := 0
	for _, val := range rus {
		if val > 500 {
			count++
		}
	}
	fmt.Println(str, "中的汉字的数量为 ", count)
}
