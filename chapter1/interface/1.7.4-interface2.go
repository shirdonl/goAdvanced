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

import "fmt"

// Message 是一个定义了通知类行为的接口
type Message interface {
	sending()
}

// 定义Programmer 及Programmer.notify方法
type Programmer struct {
	name  string
	phone string
}

func (u *Programmer) sending() {
	fmt.Printf("Sending Programmer phone to %s<%s>\n", u.name, u.phone)
}

// 定义Student及Student.message方法
type Student struct {
	name  string
	phone string
}

func (a *Student) sending() {
	fmt.Printf("Sending Student phone to %s<%s>\n", a.name, a.phone)
}

func main() {
	// 创建一个Programmer值并传给sendMessage
	bill := Programmer{"Jack", "jack@gmail.com"}
	sendMessage(&bill)

	// 创建一个Student值并传给sendMessage
	lisa := Student{"Wade", "wade@gmail.com"}
	sendMessage(&lisa)
}

// sendMessage接受一个实现了message接口的值 并发送通知
func sendMessage(n Message) {
	n.sending()
}

//Sending Programmer phone to Jack<jack@gmail.com>
//Sending Student phone to Wade<wade@gmail.com>
