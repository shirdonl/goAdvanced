package main

import (
	"fmt"
	"math/rand"
	"time"
)

//模拟请求
func simulateRequest() <-chan int32 {
	r := make(chan int32)
	go func() {
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

//乘积
func Product(a, b int32) int32 {
	return a * b
}

func main() {
	rand.Seed(time.Now().UnixNano())
	a, b := simulateRequest(), simulateRequest()
	fmt.Println(Product(<-a, <-b))
}
