package main

import "fmt"

//类型断言
func asert(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string) ///尝试转换成字符串
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println(str)
	}
}

func asert2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("string: ", t)
	case int:
		fmt.Println("int ", t)
	case bool:
		fmt.Println("bool: ", t)
	}
}

func main() {
	asert(100)
	asert("Chene")

	asert2(100)
	asert2("Chenen")

}
