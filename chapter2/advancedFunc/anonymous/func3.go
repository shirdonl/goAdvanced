package main

import "fmt"

func main() {
	f := func() string {
		return "带返回值匿名函数～"
	}
	a := f()
	fmt.Println(a) //带返回值匿名函数～
}
