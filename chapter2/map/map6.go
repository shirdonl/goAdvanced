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
	//创建map类型的切片
	slice := make([]map[int]int, 3)
	for i := range slice {
		slice[i] = make(map[int]int, 6)
		slice[i][1] = 88
		slice[i][2] = 66
	}
	fmt.Printf("Value of slice: %v\n", slice)
}

//Value of slice: [map[1:88 2:66] map[1:88 2:66] map[1:88 2:66]]
