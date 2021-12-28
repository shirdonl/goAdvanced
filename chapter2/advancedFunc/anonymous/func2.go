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
)

func main() {
	f := func(args string) {
		fmt.Println(args)
	}
	f("带参数的匿名函数～写法1")
	//或
	(func(args string) {
		fmt.Println(args)
	})("带参数的匿名函数～写法2")
	//或
	func(args string) {
		fmt.Println(args)
	}("带参数的匿名函数～写法3")
}
