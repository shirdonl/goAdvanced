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

/*
#cgo CFLAGS : -I./include
#cgo LDFLAGS: -L./lib -ltest
#include "test.h"
*/

import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.test_hello(C.CString("Shirdon"))))
}
