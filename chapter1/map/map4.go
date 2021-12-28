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
	"sort"
)

//将map按照key排序
func sortedMap(m map[int]string, f func(k int, v string)) {
	var keys []int
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		f(k, m[k])
	}
}

func main() {
	m := make(map[int]string)
	m[8] = "I"
	m[2] = "Love"
	m[6] = "Go"
	sortedMap(m, func(k int, v string) {
		fmt.Println("Key:", k, "Value:", v)
	})
}
