package main

import (
	"fmt"
	"regexp"
)

//匹配中文字符
var cnRegexp = regexp.MustCompile("^[\u4e00-\u9fa5]$")

func main() {
	str := "我爱Go"
	StrFilterGetChinese(&str)
	fmt.Println(str)
}

//获取中文字符串
func StrFilterGetChinese(src *string) {
	strNew := ""
	for _, c := range *src {
		if cnRegexp.MatchString(string(c)) {
			strNew += string(c)
		}
	}

	*src = strNew
}

//我爱
