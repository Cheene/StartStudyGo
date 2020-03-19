package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	//全部退出for 循环
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 0; j <= 10; j++ {
			if j == 2 {
				flag = true
				break
			}
			fmt.Printf("%d:%d \n", i, j)
		}
		if flag { //等价于在外层的循环处做判断
			break
		}
	}

	//打印九九
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d \t", j, i, i*j)
		}
		fmt.Println()
	}
}
