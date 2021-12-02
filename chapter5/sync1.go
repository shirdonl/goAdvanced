package main

import (
	"fmt"
	"time"
)

func main() {
	var b = 0
	for i := 0; i < 10; i++ {
		go func(idx int) {
			b += 1
			fmt.Println(b)
		}(i)
	}
	time.Sleep(time.Second)
}
