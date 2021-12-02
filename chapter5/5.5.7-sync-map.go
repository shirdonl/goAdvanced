package main

import "fmt"

func main() {
	// 创建一个int到int的映射
	m := make(map[int]int)
	// 开启goroutine
	go func() {
		// 不停地对map进行写入
		for {
			m[6] = 6
		}
	}()
	// 开启goroutine
	go func() {
		// 不停地对map进行读取
		for {
			_ = m[6]
		}
	}()
	// 无限循环, 让并发程序在后台执行
	for {
		fmt.Printf("无限循环!\n")
	}
}
