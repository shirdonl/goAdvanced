package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8888,
	})
	if err != nil {
		fmt.Println("ListenUDP error:", err)
		os.Exit(-1)
	}
	defer conn.Close()
	for {
		var buf [1024]byte
		n, addr, err := conn.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("接受消息出错：", err)
		}
		fmt.Println(addr.String(), string(buf[:n]))
		conn.WriteToUDP([]byte("收到~"), addr)
	}
}
