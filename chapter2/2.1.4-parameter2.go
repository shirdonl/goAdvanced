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

//定义1个函数，参数为 func(int, int) int
func Func(i func(int, int) int) int {
	fmt.Printf("i type: %T\n", i)
	return i(6, 9)
}

func main() {
	//定义一个匿名函数，作为Func()函数的参数
	f := func(x, y int) int {
		return x + y
	}
	fmt.Printf("f type: %T\n", f)
	//将f作为参数调用Func(f)
	fmt.Println(Func(f))
}

//f type: func(int, int) int
//i type: func(int, int) int
//15
