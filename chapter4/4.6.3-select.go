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
	"fmt"
	"os"
	"time"
)

func main() {
	//用Go的内置make()函数创建通道
	abort := make(chan struct{})
	//创建goroutine
	go func() {
		os.Stdin.Read(make([]byte, 1))
		//将结构体类型放入通道
		abort <- struct{}{}
	}()
	fmt.Println("开始倒计时......")
	tick := time.Tick(1 * time.Second) //时间间隔为１秒
	for countdown := 5; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		//释放计时器
		case <-tick:
			//什么都不做
		case <-abort:
			fmt.Println("abort...!")
			return //返回，终止
		}
	}
}
