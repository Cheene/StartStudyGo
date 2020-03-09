package main

import "fmt"

/**

 */

//全局变量声明
var (
	coins = 5000
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi",
	}

	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	//1 依次获取名字
	for _, name := range users {
		//2根据分金币的规则
		for _, c := range name {
			switch c {
			case 'e', 'E':
				distribution[name]++
				coins--
			case 'i', 'I':
				distribution[name] += 2
				coins -= 2
			case 'o', 'O':
				distribution[name] += 3
				coins -= 3

			case 'U', 'u':
				distribution[name] += 4
				coins -= 4
			}
		}
	}
	return coins
}

func main() {
	left := dispatchCoin()
	fmt.Println("还有 ", left, " Coins")

	for k, v := range distribution {
		fmt.Println("name:", k, " Coins: ", v)
	}
}
