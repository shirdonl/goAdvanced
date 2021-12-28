//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

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
