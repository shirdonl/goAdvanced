package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建map
	m := make(map[int]string)
	m[0] = "I"
	m[2] = "Go"
	m[1] = "Love"

	// 将健值按排序顺序存储在切片中
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// To perform the opertion you want
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", m[k])
	}
}

//Key: 0 Value: I
//Key: 1 Value: Love
//Key: 2 Value: Go
