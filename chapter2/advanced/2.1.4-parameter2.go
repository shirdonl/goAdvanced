package main

import "fmt"

//定义1个函数，参数为 func(int, int) int
func Func(i func(int, int) int) int {
	fmt.Printf("i type: %T\n", i)
	return i(6, 9)
}

func main() {
	//定义一个匿名函数，作为Func()函数的参数
	f := func(x, y int) int {
		return x + y
	}
	fmt.Printf("f type: %T\n", f)
	//将f作为参数调用Func(f)
	fmt.Println(Func(f))
}

//f type: func(int, int) int
//i type: func(int, int) int
//15
