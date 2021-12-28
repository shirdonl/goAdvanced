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
	// 定义局部变量
	var a int = 66
	var b int = 88

	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	// 调用函数用于交换值
	// &a 指向 a 变量的地址
	// &b 指向 b 变量的地址

	Swap(&a, &b)

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

func Swap(x *int, y *int) {
	var temp int
	temp = *x // 保存 x 地址的值
	*x = *y   // 将 y 赋值给 x[前面是a,b,后面是x,y，作者核实是否正确？][这里是对的哈，是独立的]
	*y = temp // 将 temp 赋值给 y
}

//交换前 a 的值 : 66
//交换前 b 的值 : 88
//交换后 a 的值 : 88
//交换后 b 的值 : 66
