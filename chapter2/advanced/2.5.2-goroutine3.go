package main

import (
	"fmt"
)

//Pipeline 模式
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

func Power(in <-chan int) <-chan int {
	out := make(chan int, 100)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}

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
	fmt.Println(<-Sum(Power(Generator(5))))
}

//55
