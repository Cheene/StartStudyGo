package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		for index, val := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, val)
		}
	}
	// flag.Type() 定义命令flag 参数
	name := flag.String("name", "chenen", "姓名")
	fmt.Println(name)
	flag.Parse()
	fmt.Println("flag.Args : ", flag.Args())
	fmt.Println("flag.NArg : ", flag.NArg())
	fmt.Println("flag.NFlag : ", flag.NFlag())
}
