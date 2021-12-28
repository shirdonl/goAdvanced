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
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

func main() {
	fmt.Println(CreatePassword())
}

// 生成一个随机密码
func CreatePassword() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "shirdon.liao")
	io.WriteString(h, t.String())
	password := fmt.Sprintf("%x", h.Sum(nil))
	return password
}

//372f29b807413274fe40adbc49dfd8f6
