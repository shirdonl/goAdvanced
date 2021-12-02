package main

import "fmt"

func main() {
	a := Func()
	b := a("你好～")
	c := a("你好～")
	fmt.Println(b) //世界～你好～
	fmt.Println(c) //世界～你好～你好～
}
func Func() func(string) string {
	a := "世界～"
	return func(args string) string {
		a += args
		return a
	}
}
