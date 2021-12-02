package main

import "fmt"

func main() {
	var ptr *string

	//空指针判断
	if ptr == nil { //ptr 是空指针
		fmt.Printf("ptr 的值为 : %x\n", ptr)
	}
}

//ptr 的值为 : 0
