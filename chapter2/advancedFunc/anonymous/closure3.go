package main

import "fmt"

func main() {
	arr := Func()
	arr[0]() //0xc0000b2008 0
	arr[1]() //0xc0000b2010 1
}

//定义一个闭包
func Func() []func() {
	b := make([]func(), 2, 2)
	for i := 0; i < 2; i++ {
		b[i] = (func(j int) func() {
			return func() {
				fmt.Println(&j, j)
			}
		})(i)
	}
	return b
}
