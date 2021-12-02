package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("不带参数的匿名函数～")
	}
	f()
	fmt.Printf("%T\n", f) //打印 func()
}
