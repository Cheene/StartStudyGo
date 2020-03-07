package main

import (
	"encoding/json"
	"fmt"
)

//JSON
// 1 Go ---》 JSON
//2 JSON --> Go

type person struct {
	Name string `json:"name",db:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "周琳",
		Age:  2209,
	}
	//序列化，把结构体转换到JSON
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("failed %v", err)
		return
	}
	fmt.Printf("%v\n", string(b))

	//反序列化
	str := `{"name":"嘻嘻嘻","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2) //传指针是为了能够在函数的内部修改p2
	fmt.Printf("%v\n", p2)

}
