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
)

func main() {
	fmt.Printf("%d 是否偶数：%t\n", 7, Even(7)) // 7 是否偶数：false
	fmt.Printf("%d 是否奇数：%t\n", 2, Odd(2))  // 2 是否奇数：false
	fmt.Printf("%d 是否奇数：%t\n", 3, Odd(3))  // 3 是否奇数：true
}

//判断是否是偶数
func Even(number int) bool {
	if number == 0 {
		return true
	}
	return Odd(RecursiveSign(number) - 1)
}

//判断是否是奇数
func Odd(number int) bool {
	if number == 0 {
		return false
	}
	return Even(RecursiveSign(number) - 1)
}

//递归签名
func RecursiveSign(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

//7 是否偶数：false
//2 是否奇数：false
//3 是否奇数：true
