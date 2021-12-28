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
	"reflect"
)

func main() {
	a := make([]int, 6)
	for i := 0; i < 6; i++ {
		a[i] = i + 2
	}
	index := arrayPosition(a, 6)
	fmt.Println(index)
}

//查找一个元素在切片中的位置
func arrayPosition(arr interface{}, d interface{}) int {
	array := reflect.ValueOf(arr)
	for i := 0; i < array.Len(); i++ {
		v := array.Index(i)
		if v.Interface() == d {
			return i
		}
	}
	return -1
}

//4
