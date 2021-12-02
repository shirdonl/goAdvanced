package main

import "fmt"

func Hello(c chan string) {
	name := <-c // 从通道获取数据
	fmt.Println("Hello", name)
}

func main() {

	// 声明一个字符串类型的变量
	ch := make(chan string)

	// 开启一个 goroutine
	go Hello(ch)

	// 发送数据到通道 ch
	ch <- "World"

	//关闭通道
	close(ch)
}
