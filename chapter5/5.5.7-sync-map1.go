package main

import (
	"fmt"
	"sync"
)

func main() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("Jack", 90)
	scene.Store("Barry", 99)
	scene.Store("ShirDon", 100)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("Barry"))
	// 根据键删除对应的键值对
	scene.Delete("Jack")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

//99 true
//iterate: Barry 99
//iterate: ShirDon 100
