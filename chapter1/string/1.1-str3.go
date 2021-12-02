package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ILoveGoAdvancedProgramming!"
	content := GetSubStrIndex(str, "Go")
	fmt.Println(content)

	str1 := "我爱高级编程!"
	content1 := GetSubString(str1, 2, 5)
	fmt.Println(content1)
}

//获取子串的位置
func GetSubStrIndex(str, substr string) int {
	// 子串在字符串中的字节位置
	result := strings.Index(str, substr)
	if result >= 0 {
		// 获得子串之前的字符串并转换成[]byte
		prefix := []byte(str)[0:result]
		// 将子串之前的字符串转换成[]rune
		res := []rune(string(prefix))
		// 获得子串之前的字符串的长度，即子串在字符串的字符位置
		result = len(res)
	}

	return result
}

//获取中文字符串
func GetSubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}

	// 返回子串
	return string(rs[begin:end])
}
