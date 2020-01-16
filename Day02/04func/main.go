package main

import "fmt"

// 明明返回值就相当于在函数中声明了一个int类型的变量
func add(a int, b int) (ret int) {
	return a + b
}

// 无参无返回值
func f2() {
	fmt.Println("f2")
}

// 有返回值，
func f3() int {
	return 3
}

// 多返回值
func f5() (int, string) {
	return 2019, "chenene"
}

//类型的简写 x,y Y可以不传，可以只传一个，也可以多传几个。但只能放在最后
func f6(x, y string) string {
	return x + "" + y
}

// 可变长参数
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	r := add(1, 3)
	fmt.Println(r)
	f2()
	fmt.Println(f3())

	x, y := f5()
	fmt.Println(x, y)

	strAdd := f6("aa", "bb")
	fmt.Println(strAdd)

	f7("No.1")
	f7("No.2", 3, 4, 43534, 5634, 534, 534)
}
