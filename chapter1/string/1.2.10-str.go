package main

import (
	"fmt"
	"strings"
)

func main() {
	// ,|/ 都是分隔符
	fn := func(c rune) bool {
		return strings.ContainsRune(",|/", c)
	}
	str := strings.FieldsFunc("Python,Jquery|JavaScript,Go,C++/C", fn)
	fmt.Println(str)
}

//[Python Jquery JavaScript Go C++ C]
