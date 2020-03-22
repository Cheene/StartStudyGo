package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Printf("Err for dail 127.0.0.1:8090, %#v", err)
		return
	}
	defer conn.Close()
	//发送
	for i := 0; i < 20; i++ {
		msg := `Hello Chenen.None of your bussiness`
		conn.Write([]byte(msg))
	}
}
