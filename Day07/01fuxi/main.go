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
	str := `{"name":"chenene","age":200}`
	var stu1 person
	json.Unmarshal([]byte(str), &stu1)
	fmt.Println(stu1)
}
