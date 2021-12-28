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

func main() {
	a := Func()
	b := a("你好～")
	c := a("你好～")
	fmt.Println(b) //世界～你好～
	fmt.Println(c) //世界～你好～你好～
}
func Func() func(string) string {
	a := "世界～"
	return func(args string) string {
		a += args
		return a
	}
}
