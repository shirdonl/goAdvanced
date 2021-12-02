package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " Go Advanced  "
	str1 := strings.TrimSpace(str)
	str2 := strings.Trim(str, " ")
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(str2)
}

// Go Advanced
//Go Advanced
//Go Advanced
