package main

import (
	"fmt"
)

func main() {
	var arr = []interface{}{1, 6, 8}

	fmt.Println(arr)
	fmt.Println(arr...)
}

//[1 6 8]
//1 6 8
