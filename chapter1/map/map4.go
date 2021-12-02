package main

import (
	"fmt"
	"sort"
)

//将map按照key排序
func sortedMap(m map[int]string, f func(k int, v string)) {
	var keys []int
	for k, _ := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		f(k, m[k])
	}
}

func main() {
	m := make(map[int]string)
	m[8] = "I"
	m[2] = "Love"
	m[6] = "Go"
	sortedMap(m, func(k int, v string) {
		fmt.Println("Key:", k, "Value:", v)
	})
}
