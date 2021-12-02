package main

import "fmt"

//检查字符串是否在切片中
func Exist(target string, array []string) bool {
	for _, element := range array {
		if target == element {
			return true
		}
	}

	return false
}

func main() {
	nameList := []string{"Barry", "Shirdon", "Jack"}
	str1 := "Barry"
	str2 := "Go"
	result := Exist(str1, nameList)
	fmt.Println("Barry 是否在 nameList 中：", result)
	result = Exist(str2, nameList)
	fmt.Println("Go是否在 nameList 中：", result)
}

//Barry 是否在 nameList 中： true
//Go是否在 nameList 中： false
