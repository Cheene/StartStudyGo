package main

import (
	"../split"
	"fmt"
)

func main() {
	ret := split.Split("babcdfbef", "b")
	fmt.Printf("%#v\n", ret)

}
