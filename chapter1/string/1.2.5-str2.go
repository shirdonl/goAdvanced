package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	test := "I,Love,Go"
	str := test
	keywordSlice := strings.Split(test, ",")
	for _, v := range keywordSlice {
		reg := regexp.MustCompile("(?i)" + v)
		str = reg.ReplaceAllString(str, strings.ToUpper(v))
		fmt.Println(str)
	}
}

//I,Love,Go
//I,LOVE,Go
//I,LOVE,GO
