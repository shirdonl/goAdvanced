package main

import (
	"fmt"
)

func main() {
	var array = []int{1, -2, 88, 66, 16, 68}
	maxValue := array[0]
	maxValueIndex := 0
	for i := 0; i < len(array); i++ {
		//比较元素的大小，如果发现有更大的数，则进行交换
		if maxValue < array[i] {
			maxValue = array[i]
			maxValueIndex = i
		}
	}
	fmt.Printf("maxValue=%v maxValueIndex=%v \n", maxValue, maxValueIndex)
}
