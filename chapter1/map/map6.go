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
	// 创建
	m := map[string]string{
		"a": "value_a",
		"b": "value_b",
	}

	for k, v := range m {
		fmt.Printf("k:[%v].v:[%v]\n", k, v) // 输出k,v值
	}
}

//k:[a].v:[value_a]
//k:[b].v:[value_b]
