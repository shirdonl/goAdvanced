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

import (
	"fmt"
	"gitee.com/shirdonl/goAdvanced/chapter2/array"
)

func main() {
	//定义一个二维数组
	nums := [][]int{{1, 9, 5}, {2, 3, 6}, {3, 6, 9}, {1, 8, 3}}

	firstIndex := 2 //按第二列排序

	result := array.ArraySort(nums, firstIndex-1)

	fmt.Println(result)
}
