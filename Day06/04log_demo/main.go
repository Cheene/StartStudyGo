package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("./some.log", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Open failed ", err)
		return
	}
	log.SetOutput(file)
	for {
		v := 100
		log.Printf("测试%v", v)
		time.Sleep(time.Second * 2)
	}
}
