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

	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 5; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 1; j <= 5; j++ {
			fmt.Printf("j: %d\n", j)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- break 标签 ----")
breakLabel:
	for i := 1; i <= 5; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 1; j <= 5; j++ {
			fmt.Printf("j: %d\n", j)
			continue breakLabel
		}
	}
}
