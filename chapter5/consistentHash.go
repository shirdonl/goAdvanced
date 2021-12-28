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
	"github.com/golang/groupcache/consistenthash"
)

func main() {
	// 构造一个 consistenthash 对象，每个节点在 Hash 环上都一共有4个虚拟节点。
	hash := consistenthash.New(4, nil)

	// 添加节点
	hash.Add(
		"10.0.0.1:8080",
		"10.0.0.1:8081",
		"10.0.0.1:8082",
		"10.0.0.1:8083",
	)

	// 根据 key 获取其对应的节点
	//Get 获取散列中与提供的键最接近的项目
	node := hash.Get("10.0.0.1")
	fmt.Println(node)
}
