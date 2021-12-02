package main

import (
	"fmt"
)

//工序1：数组生成器，生成一个数组
func Generator(max int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for i := 1; i <= max; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

//工序2：求一个整数的平方
func Square(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

//工序3：求和
func Sum(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		var Sum int
		for v := range in {
			Sum += v
		}
		out <- Sum
		close(out)
	}()
	return out
}

func main() {
	// [1, 2, 3, 4, 5]
	//工序1：生成数组
	arr := Generator(5)
	//工序2：求数组每一个元素的平方
	squ := Square(arr)
	//工序3：求所有元素的和
	sum := <-Sum(squ)
	//打印
	fmt.Println(sum)
}

//55
