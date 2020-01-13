package main

import "fmt"

func main() {
	f1 := 1.2345 // 默认情况下 float是64位
	fmt.Printf("f1: %T\n", f1)
	f2 := float32(1.2345)
	fmt.Printf("f2: %T\n", f2)
	// float32类型的不能直接赋值给float64
	//f1 = f2
	//复数
	var c1 complex64 // 实部和虚部各32位
	c1 = 1 + 2i
	c2 := 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

}
