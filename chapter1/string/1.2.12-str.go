package main

import (
	"fmt"
	"strings"
)

func main() {
	fn := func(c rune) bool {
		return strings.ContainsRune(",|/", c)
	}
	res := strings.TrimFunc("|/Shirdon Liao,/", fn)
	fmt.Println(res)
}

//Shirdon Liao
