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

import "fmt"

var counter = func(n int) chan<- chan<- int {
	//定义一个chan<- int类型的chan
	requests := make(chan chan<- int)
	go func() {
		for request := range requests {
			if request == nil {
				n++ // 递增计数
			} else {
				request <- n // 返回当前计数
			}
		}
	}()
	return requests // 隐式转换到类型chan<- (chan<- int)
}(0)

func main() {
	increase100 := func(done chan<- struct{}) {
		for i := 0; i < 100; i++ {
			counter <- nil
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase100(done)
	go increase100(done)
	<-done
	<-done

	request := make(chan int, 1)
	counter <- request
	fmt.Println(<-request) // 200
}
