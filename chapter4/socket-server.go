package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Log(msg string, err error) {
	fmt.Println(msg, err)
	os.Exit(-1)
}

func Work(sock net.Conn) {
	for {
		var buf [1024]byte
		n, err := sock.Read(buf[:])
		if err != nil {
			Log("读数据出错：", err)
		}
		fmt.Printf("%v说：%v", sock.RemoteAddr().String(), string(buf[:n]))

		var str string
		reader := bufio.NewReader(os.Stdin)
		str, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("发送失败：", err)
			os.Exit(-1)
		}
		sock.Write([]byte(str))
	}
}

func main() {
	listenSock, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		Log("监听失败：", err)
	}
	for {
		connectSock, err := listenSock.Accept()
		defer connectSock.Close()
		if err != nil {
			Log("接受请求失败：", err)
		}
		go Work(connectSock)
	}
}
