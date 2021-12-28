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
	var str string = "Barry" // 声明实际变量
	var name *string         // 声明指针变量

	name = &str // 指针变量的存储地址

	fmt.Printf("str 变量的地址是: %x\n", &str)
	// 指针变量的存储地址
	fmt.Printf("name 变量储存的指针地址: %x\n", name)
	// 使用指针访问值
	fmt.Printf("*name 变量的值: %s\n", *name)
}

//str 变量的地址是: c000010200
//name 变量储存的指针地址: c000010200
//*name 变量的值: Barry
