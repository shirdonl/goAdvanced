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
	"log"
	"net/http"
)

func hiWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello,Web World!")
}

func main() {
	http.HandleFunc("/hi", hiWorld)
	if err := http.ListenAndServe(":8085", nil); err != nil {
		log.Fatal(err)
	}
}
