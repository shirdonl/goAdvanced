package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr map[int]int
	arr = make(map[int]int, 5)
	arr[0] = 88
	arr[1] = 66
	arr[2] = 99
	//定义一个切片
	var b []int
	fmt.Println("排序前的值如下：")
	//注意map是无序的
	for k, v := range arr {
		fmt.Println(k, v)
		b = append(b, v)
	}

	sort.Ints(b)
	fmt.Println("排序后的值如下：")
	for k, v := range b {
		fmt.Println(k, v)
	}
}

//排序前的值如下：
//0 88
//1 66
//2 99
//排序后的值如下：
//0 66
//1 88
//2 99
