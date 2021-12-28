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
	"strings"
)

func main() {
	s := "I_Love_Go_Advanced"
	res := strings.Split(s, "_")
	for i := range res {
		fmt.Println(res[i])
	}
	res1 := strings.SplitN(s, "_", 2)
	for i := range res1 {
		fmt.Println(res1[i])
	}
	res2 := strings.SplitAfter(s, "_")
	for i := range res2 {
		fmt.Println(res2[i])
	}
	res3 := strings.SplitAfterN(s, "_", 2)
	for i := range res3 {
		fmt.Println(res3[i])
	}
}
