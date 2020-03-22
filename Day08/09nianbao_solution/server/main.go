package main

import (
	"../proto"
	"bufio"
	"fmt"
	"io"
	"net"
)

//处理并发
func processConn(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		//已经读到文件的结尾处
		if err == io.EOF {
			break
		}
		fmt.Println(msg)
	}

}

//tcp  server
func main() {
	//1 启动服务
	listener, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Printf("Tcp Start Server Error: %#v\n", err)
		return
	}
	defer listener.Close()
	//2 等待连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed: ", err)
		}
		//3 成功后与客户端通信
		go processConn(conn)
	}
}
