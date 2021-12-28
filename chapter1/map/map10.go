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

type frequency struct {
	char string
	fre  int
}

//计算字符串的字符出现频率，并按照频率降序排序
func frequencies(s string) []frequency {
	m := make(map[string]int)
	for _, r := range s {
		m[string(r)]++
	}
	a := make([]frequency, 0, len(m))
	for c, f := range m {
		a = append(a, frequency{char: c, fre: f})
	}
	sort.Slice(a, func(i, j int) bool { return a[i].fre > a[j].fre })
	return a
}

func main() {
	str := "hiilovegogogo"
	fmt.Println(str)
	f := frequencies(str)
	fmt.Println(f)
}

//[{o 4} {g 3} {i 2} {l 1} {v 1} {e 1} {h 1}]
