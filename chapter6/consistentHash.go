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
