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
	slice := make([]int, 0)
	slice = append(slice, 1, 6, 8)
	var I interface{} = slice
	if res, ok := I.([]int); ok {
		fmt.Println(res) //[1 6 8]
		fmt.Println(ok)
	}
}

func Len(array interface{}) int {
	var length int //数组的长度
	if array == nil {
		length = 0
	}
	switch array.(type) {
	case []int:
		length = len(array.([]int))
	case []string:
		length = len(array.([]string))
	case []float32:
		length = len(array.([]float32))

	default:
		length = 0
	}
	fmt.Println(length)

	return length
}
