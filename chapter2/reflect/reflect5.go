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
