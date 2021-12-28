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
	var a interface{} = func(a int) string {
		return fmt.Sprintf("d:%d", a)
	}
	switch b := a.(type) { // 局部变量b是类型转换后的结果
	case nil:
		println("nil")
	case *int:
		println(*b)
	case func(int) string:
		println(b(88))
	case fmt.Stringer:
		fmt.Println(b)
	default:
		println("unknown")
	}
}

//d:88
