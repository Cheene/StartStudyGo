package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:name`
	Age  int    `json:age`
}

func main() {
	str := `{"name":"chen","age":100}`
	var p person
	json.Unmarshal([]byte(str), &p) //反射的操作
	fmt.Println(p.Name, p.Age)
}
