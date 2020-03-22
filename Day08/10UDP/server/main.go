package main

import (
	"fmt"
	"net"
	"strings"
)

//UDP server

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 12000,
		Zone: "",
	})
	if err != nil {
		fmt.Printf("Err Start Server :%#v", err)
		return
	}
	defer conn.Close()
	//收发数据，无连接的地址
	for {
		var data [1024]byte
		//数组转换成切片
		n, addr, err := conn.ReadFromUDP(data[:])
		fmt.Println(addr)
		if err != nil {
			fmt.Println("read udp failed,err: ", err)
			continue
		}
		//数据打印到终端
		fmt.Printf("data:%v,addr:#v,count:%v\n", string(data[:]), addr, n)
		//数据返回
		reply := strings.ToUpper(string(data[:n]))
		_, err = conn.WriteToUDP([]byte(reply), addr)
		if err != nil {
			fmt.Println("write to udp failed,err: ", err)
			continue
		}
	}

}
