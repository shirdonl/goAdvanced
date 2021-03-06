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
	"gitee.com/shirdonl/goAdvanced/chapter3/distribute"
	"os"
	"unsafe"
)

// Master运行方式形如：go run 4.8-distribute-crawer.go master 127.0.0.1:8806
// Worker运行方式形如：go run main.go 4.8-distribute-crawer.go 127.0.0.1:8806 127.0.0.1:8808 &
func main() {
	if len(os.Args) < 2 {
		sz := 100
		p := make([]int, sz)
		p[2] = 5
		fmt.Printf("%s: 使用日志请查看文件： %d\n", os.Args[0], *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&p[0])) + 2*unsafe.Sizeof(p[0]))))
		return
	}
	switch os.Args[1] {
	case "master":
		if len(os.Args) == 3 {
			distribute.RunMaster(os.Args[2])
		}
	case "worker":
		if len(os.Args) == 4 {
			distribute.RunWorker(os.Args[2], os.Args[3])
		}
	}
}
