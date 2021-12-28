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
	f1, f2 := MultiFunc(6, 8)
	fmt.Println(f1(1)) //7
	fmt.Println(f2())  //28
}

//返回多个匿名函数
func MultiFunc(a, b int) (func(int) int, func() int) {
	f1 := func(c int) int {
		return (a + b) * c / 2
	}

	f2 := func() int {
		return 2 * (a + b)
	}
	return f1, f2
}
