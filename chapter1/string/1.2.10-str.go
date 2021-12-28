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
	// ,|/ 都是分隔符
	fn := func(c rune) bool {
		return strings.ContainsRune(",|/", c)
	}
	str := strings.FieldsFunc("Python,Jquery|JavaScript,Go,C++/C", fn)
	fmt.Println(str)
}

//[Python Jquery JavaScript Go C++ C]
