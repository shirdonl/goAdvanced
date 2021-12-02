package main

import (
	"fmt"
	"time"
)

func Count(c chan<- int) {
	for i := 0; i < 6; i++ {
		time.Sleep(2 * time.Second)
		c <- i
	}
	close(c)
}

func main() {
	msg := "开启main()函数"
	fmt.Println(msg)
	ch := make(chan int)
	go Count(ch)
	for count := range ch {
		fmt.Println("count:", count)
	}
}
