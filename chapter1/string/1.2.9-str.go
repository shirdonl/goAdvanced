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
)

//匹配中文字符
var cnRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

func main() {
	str := "我爱Go"
	StrFilterGetChinese(&str)
	fmt.Println(str)
}

//获取中文字符串
func StrFilterGetChinese(src *string) {
	strNew := ""
	for _, c := range *src {
		if cnRegexp.MatchString(string(c)) {
			strNew += string(c)
		}
	}

	*src = strNew
}

//我爱
