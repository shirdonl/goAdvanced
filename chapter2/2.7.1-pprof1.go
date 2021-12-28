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
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// 性能分析
	go func() {
		log.Println(http.ListenAndServe(":8802", nil))
	}()

	// 实际业务代码
	for {
		Sum("This is a test")
	}
}

func Sum(str string) string {
	data := []byte(str)
	sData := string(data)
	var sum = 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	return sData
}
