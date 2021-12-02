package main

import (
	"fmt"
	"gitee.com/shirdonl/goAdvanced/chapter2/array"
)

func main() {
	//定义一个二维切片
	nums := [][]int{{1, 9, 5}, {2, 3, 6}, {3, 6, 9}, {1, 8, 3}}

	firstIndex := 2 //按第二列排序

	result := array.ArraySort(nums, firstIndex-1)

	fmt.Println(result)
}
