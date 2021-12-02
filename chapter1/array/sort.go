package array

import (
	"fmt"
	"sort"
)

//按指定规则对numArray进行排序
func ArraySort(numArray [][]int, firstIndex int) [][]int {

	//检查
	if len(numArray) <= 1 {
		return numArray
	}

	if firstIndex < 0 || firstIndex > len(numArray[0])-1 {
		fmt.Println("Warning: Param firstIndex should between 0 and len(numArray)-1. The original array is returned.")
		return numArray
	}

	//排序
	mIntArray := &IntArray{numArray, firstIndex}
	sort.Sort(mIntArray)
	return mIntArray.mArr
}

type IntArray struct {
	mArr       [][]int
	firstIndex int
}

//IntArray实现sort.Interface接口
func (arr *IntArray) Len() int {
	return len(arr.mArr)
}

func (arr *IntArray) Swap(i, j int) {
	arr.mArr[i], arr.mArr[j] = arr.mArr[j], arr.mArr[i]
}

func (arr *IntArray) Less(i, j int) bool {
	arr1 := arr.mArr[i]
	arr2 := arr.mArr[j]
	for index := arr.firstIndex; index < len(arr1); index++ {
		if arr1[index] < arr2[index] {
			return true
		} else if arr1[index] > arr2[index] {
			return false
		}
	}

	return i < j
}
