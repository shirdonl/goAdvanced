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
	var mArray [1][2][3]int
	fmt.Println(mArray)
	fmt.Println(mArray[0])
	fmt.Println(mArray[0][0])
	fmt.Println(mArray[0][0][0])
}

//[[[0 0 0] [0 0 0]]]
//[[0 0 0] [0 0 0]]
//[0 0 0]
//0
