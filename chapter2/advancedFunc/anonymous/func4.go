package main

import "fmt"

func main() {
	f1, f2 := MultiFunc(6, 8)
	fmt.Println(f1(1)) //7
	fmt.Println(f2())  //28
}

//返回多个匿名函数
func MultiFunc(a, b int) (func(int) int, func() int) {
	f1 := func(c int) int {
		return (a + b) * c / 2
	}

	f2 := func() int {
		return 2 * (a + b)
	}
	return f1, f2
}
