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
	var money float32 = 66.66
	fmt.Println("指针原来的值是：", money)

	// 通过reflect.ValueOf获取money中的reflect.Value
	// 注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&money)
	newValue := pointer.Elem()

	fmt.Println("指针的类型是：", newValue.Type())
	fmt.Println("指针是否可设置：", newValue.CanSet())

	// 重新赋值
	newValue.SetFloat(88.88)
	fmt.Println("指针的新值是：", money)
}

//指针原来的值是： 66.66
//指针的类型是： float32
//指针是否可设置： true
//指针的新值是： 88.88
