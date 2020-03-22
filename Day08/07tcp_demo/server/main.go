package main

import (
	"fmt"
	"net"
)

//处理并发
func processConn(conn net.Conn) {
	var temp [128]byte
	for {
		n, err := conn.Read(temp[:])
		if err != nil {
			fmt.Println("read failed: ", err)
		}
		fmt.Println(string(temp[:n]))
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
