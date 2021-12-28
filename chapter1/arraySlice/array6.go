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

import (
	"fmt"
)

func main() {
	array := []int{1, 6, 6, 8}
	res := removeDuplicates(array)
	fmt.Println(res)
}

//删除重复元素
func removeDuplicates(array []int) []int {
	//如果是空切片，则返回nil
	if len(array) == 0 {
		return nil
	}
	//用两个标记来比较相邻位置的值
	//如果一样，则继续
	//如果不一样，则把right指向的值赋值给left下一位
	left, right := 0, 1
	for ; right < len(array); right++ {
		if array[left] == array[right] {
			continue
		}
		left++
		array[left] = array[right]
	}
	return array[:left+1]
}
