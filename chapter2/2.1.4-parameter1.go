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
	/* 定义局部变量 */
	num1 := 6
	num2 := 8
	fmt.Printf("交换前num1的值为 : %d\n", num1)
	fmt.Printf("交换前num2的值为 : %d\n", num2)
	/* 通过调用函数来交换值 */
	exchange(num1, num2)
	fmt.Printf("交换后num1的值 : %d\n", num1)
	fmt.Printf("交换后num2的值 : %d\n", num2)
}

/* 定义相互交换值的函数 */
func exchange(x, y int) int {
	var tmp int
	tmp = x /* 保存 x 的值 */
	x = y   /* 将 y 值赋给 x */
	y = tmp /* 将 temp 值赋给 y*/
	return tmp
}

//交换前num1的值为 : 6
//交换前num2的值为 : 8
//交换后num1的值 : 6
//交换后num2的值 : 8
