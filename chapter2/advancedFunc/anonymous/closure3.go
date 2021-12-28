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
	arr := Func()
	arr[0]() //0xc0000b2008 0
	arr[1]() //0xc0000b2010 1
}

//定义一个闭包
func Func() []func() {
	b := make([]func(), 2, 2)
	for i := 0; i < 2; i++ {
		b[i] = (func(j int) func() {
			return func() {
				fmt.Println(&j, j)
			}
		})(i)
	}
	return b
}
