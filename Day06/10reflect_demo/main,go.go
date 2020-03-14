package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string `json:name`
	Age  int    `json:age`
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind() //值的种类
	switch k {
	case reflect.Int:
		fmt.Printf("Int : %v\n", int(v.Int()))
	case reflect.String:
		fmt.Printf("String : %v\n", string(v.Int()))
	}
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
	fmt.Printf("type name:%v, type kind:%v\n", v.Name(), v.Kind()) //类型的种类
}

func main() {
	var a float32 = 2
	reflectType(a)
	var b int16 = 2
	reflectType(b)

	reflectType(person{})
	var qq int = 3
	reflectValue(qq)
}
