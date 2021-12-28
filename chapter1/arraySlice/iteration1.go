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
	// 创建一个整型切片，并赋值
	array := []int{1, 6, 8}
	// 迭代每一个元素，并打印其值
	for k, v := range array {
		fmt.Printf("index:%d value:%d  element-address:%X \n", k, v, &array[k])
	}
	fmt.Println("\n使用匿名变量（下划线）来忽略索引值：")
	//使用空白标识符（下划线）来忽略索引值
	for _, v := range array {
		fmt.Printf(" value:%d   \n", v)
	}
	fmt.Println("\n使用 for 循环对切片进行迭代：")
	//使用传统的 for 循环对切片进行迭代
	for i := 0; i < len(array); i++ {
		fmt.Printf(" value:%d   \n", array[i])
	}
}
