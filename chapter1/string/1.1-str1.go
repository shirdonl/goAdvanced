package main

import (
	"fmt"
	"strings"
)

func stringLength(str string) int {
	var Length int
	var str1 string
	left := 0
	right := 0
	str1 = str[left:right]
	for ; right < len(str); right++ {
		//IndexByte函数的功能是检查字节c在str中第一次出现的位置索引;
		//如果s中第一次出现则返回-1，第二次出现则返回第二次出现的位置索
		index := strings.IndexByte(str1, str[right])
		//当某个字符出现第二次的时候返回当前索引加1的位置
		if index != -1 {
			left = index + 1
			fmt.Println(str1)
		}
		//字符串长度截取，left起始位置，right截止位置
		str1 = str[left : right+1]
		//获取到的长度如果大于先前的长度，就替换直到获取到最长字符串
		if len(str1) > Length {
			Length = len(str1)
		}
	}
	return Length
}

func main() {
	s := stringLength("ilikegoadvanced")
	fmt.Print("\n", s, "\n")
}
