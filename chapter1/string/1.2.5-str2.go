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
	"regexp"
	"strings"
)

func main() {
	test := "I,Love,Go"
	str := test
	keywordSlice := strings.Split(test, ",")
	for _, v := range keywordSlice {
		reg := regexp.MustCompile("(?i)" + v)
		str = reg.ReplaceAllString(str, strings.ToUpper(v))
		fmt.Println(str)
	}
}

//I,Love,Go
//I,LOVE,Go
//I,LOVE,GO
