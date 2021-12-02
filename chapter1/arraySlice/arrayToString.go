package main

import "fmt"

func main() {
	array := make([]string, 5)
	array[0] = "ILoveGo"
	string := array[0]               //直接将该数组的一个元素，赋值给字符串
	fmt.Printf("====>:%s\n", string) //====>:ILoveGo

	str := arrayToString(array)
	fmt.Printf("====>:%s\n", str)
}

func arrayToString(arr []string) string {
	var result string
	for _, i := range arr { //遍历数组中所有元素追加成string
		result += i
	}
	return result
}
