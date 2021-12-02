package main

import "fmt"

func main() {
	arr := Func()
	arr[0]() //0xc000014070 2
	arr[1]() //0xc000014070 2
}

//定义一个闭包
func Func() []func() {
	//声明一个匿名函数类型的数组
	b := make([]func(), 2, 2)
	for i := 0; i < 2; i++ {
		b[i] = func() {
			fmt.Println(&i, i)
		}
	}
	return b
}
