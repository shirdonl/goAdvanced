package main

import (
	"fmt"
)

func main() {
	f := func(args string) {
		fmt.Println(args)
	}
	f("带参数的匿名函数～写法1")
	//或
	(func(args string) {
		fmt.Println(args)
	})("带参数的匿名函数～写法2")
	//或
	func(args string) {
		fmt.Println(args)
	}("带参数的匿名函数～写法3")
}
