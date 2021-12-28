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
	//定义字符串
	str := "123456789abc"
	//调用反转函数
	strRev := Reversal(str)
	//打印
	fmt.Println(str)
	fmt.Println(strRev)
}

//反转字符串
func Reversal(a string) (re string) {
	//将字符串转成rune数组
	b := []rune(a)
	//遍历
	for i := 0; i < len(b)/2; i++ {
		//交换
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}
	//转换为字符串类型
	re = string(b)
	return
}

//123456789abc
//cba987654321
