//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import "fmt"

func main() {
	var mArray [3][2]int
	mArray[0] = [2]int{6, 8}
	fmt.Println(mArray)
	mArray[1] = [2]int{66, 88}
	fmt.Println(mArray)
}

//[[6 8] [0 0] [0 0]]
//[[6 8] [66 88] [0 0]]
