package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	server, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("拨号失败：", err)
		os.Exit(-1)
	}
	defer server.Close()
	for {
		var str string
		var buf [1024]byte
		reader := bufio.NewReader(os.Stdin)
		str, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("发送失败：", err)
			os.Exit(-1)
		}
		server.Write([]byte(str))
		var n int
		n, err = server.Read(buf[:])
		if err != nil {
			fmt.Println("读取失败：", err)
			os.Exit(-1)
		}
		fmt.Printf("%v说：%v", server.RemoteAddr().String(), string(buf[:n]))
	}
}
