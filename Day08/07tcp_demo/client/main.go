package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

//tcp client

func main() {
	//1 与server端建立连接‘
	conn, err := net.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Printf("dail 127.0.0.1:8090 failed, %#v\n", err)
		return
	}
	//2 发送数据
	reader := bufio.NewReader(os.Stdin)
	//if len(os.Args) < 2 {
	//	msg =  "chene"
	//} else {
	//	msg = os.Args[1]
	//}
	for {
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		conn.Write([]byte(msg))
	}
	//3 结束通讯
	conn.Close()
}
