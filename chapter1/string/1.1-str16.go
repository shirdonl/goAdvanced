package main

import (
	"fmt"
)

func main() {
	str := "go"
	r := StringToBin(str)
	fmt.Println(r)
}

func StringToBin(s string) (binString string) {
	for _, c := range s {
		binString = fmt.Sprintf("%s%b", binString, c)
	}
	return
}

//11001111101111
