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
	"time"
)

func main() {
	fmt.Println("开始倒计时......")
	tick := time.Tick(1 * time.Second) //时间间隔为１秒
	for countdown := 5; countdown > 0; countdown-- {
		fmt.Println(countdown)
		//释放计时器
		<-tick
	}
}

//开始倒计时......
//5
//4
//3
//2
//1
