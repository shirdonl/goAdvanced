package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8888,
	})
	if err != nil {
		fmt.Println("拨号失败：", err)
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		str = strings.Trim(str, "\r\n")
		if strings.ToUpper(str) == "Q" {
			fmt.Println("退出聊天系统～")
			return
		}
		con.Write([]byte(str))

		var buf [1024]byte
		n, _, _ := con.ReadFromUDP(buf[:])
		fmt.Println(con.RemoteAddr().String(), string(buf[:n]))
	}
}
