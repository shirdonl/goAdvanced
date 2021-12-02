package main

import "fmt"

func main() {
	//声明一个数字常量
	const MAX int = 3
	//声明一个数组
	arr := []int{66, 88, 99}
	var i int
	//声明整型指针数组
	var ptr [MAX]*int

	for i = 0; i < MAX; i++ {
		ptr[i] = &arr[i] // 整数地址赋值给指针数组
	}

	//循环打印
	for i = 0; i < MAX; i++ {
		fmt.Printf("arr[%d] = %d\n", i, *ptr[i])
	}
}

//arr[0] = 66
//arr[1] = 88
//arr[2] = 99
