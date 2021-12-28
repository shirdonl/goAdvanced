//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
	"math/rand"
	"time"
)

//模拟请求
func simulateRequest() <-chan int32 {
	r := make(chan int32)
	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

//乘积
func Product(a, b int32) int32 {
	return a * b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := simulateRequest(), simulateRequest()
	fmt.Println(Product(<-a, <-b))
}
