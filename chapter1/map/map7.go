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
	// 创建map
	m := map[string]string{
		"a": "value_a",
		"b": "value_b",
	}

	var sli []*string
	for k, v := range m {
		//k一直使用同一块内存，v也一样
		fmt.Printf("k:[%p].v:[%p]\n", &k, &v)
		sli = append(sli, &v) // 对v取地址的值
	}
	// 输出
	//k:[0xc000010200].v:[0xc000010210]
	//k:[0xc000010200].v:[0xc000010210]

	for _, b := range sli {
		fmt.Println(*b)
	}
	// 输出
	// value_b
	// value_b
}
