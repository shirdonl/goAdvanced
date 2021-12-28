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
	"reflect"
)

func main() {
	var money float32 = 66.68

	pointer := reflect.ValueOf(&money)
	value := reflect.ValueOf(money)

	convertPointer := pointer.Interface().(*float32)
	convertValue := value.Interface().(float32)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

//0xc0000b2004
//66.68
