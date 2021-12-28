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
	var str string
	var ptr *string
	var pptr **string

	str = "Go Advanced"

	//指针 ptr 地址
	ptr = &str

	// 指向指针 ptr 地址
	pptr = &ptr

	// 获取 pptr 的值
	fmt.Printf("变量 str = %s\n", str)
	fmt.Printf("指针变量 *ptr = %s\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %s\n", **pptr)
}

//变量 str = Go Advanced
//指针变量 *ptr = Go Advanced
//指向指针的指针变量 **pptr = Go Advanced
