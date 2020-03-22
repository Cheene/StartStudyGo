package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 12000,
		Zone: "",
	})

	if err != nil {
		fmt.Printf("server connect faield: %#v", err)
		return
	}

	var reply [1024]byte
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("请输入内容: ")
		msg, _ := reader.ReadString('\n')
		socket.Write([]byte(msg))
		n, _, err := socket.ReadFromUDP(reply[:])
		if err != nil {
			fmt.Println("receiver data error: ", err)
			return
		}
		fmt.Printf("data:%v,count:%v", string(reply[:]), n)
	}
}
