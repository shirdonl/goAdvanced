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
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	data := []byte("hello,sha1")
	fmt.Printf("%x \n", sha1.Sum(data))

	h := sha1.New()
	io.WriteString(h, "hello,sha1")
	fmt.Printf("%x \n", h.Sum(nil))
}

//73eb2c728f810753c9537ab9b876c9c0305255f5
//73eb2c728f810753c9537ab9b876c9c0305255f5
