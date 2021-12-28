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

import "gitee.com/shirdonl/goAdvanced/chapter4/sync/pkg"

func main() {
	var ch = make(chan string, 2)
	// 结束标志
	var finished = make(chan bool)

	go pkg.TaskProducer(ch)
	go pkg.TaskConsumer(ch, finished)

	<-finished
}

//Task7:9,8
//Task6:1,7
//Task3:5,0
//Task4:6,4
//Task2:2,7
//Task5:1,9
//Task1:0,8
//Task8:4,1
