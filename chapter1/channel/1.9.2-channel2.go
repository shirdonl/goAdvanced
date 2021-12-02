package main

import (
	"fmt"
	"time"
)

func TimeLong(d time.Duration) <-chan struct{} {
	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(d)
		ch <- struct{}{}
	}()
	return ch
}

func main() {
	fmt.Println("你好～")
	<-TimeLong(time.Second)
	fmt.Println("过1秒后继续显示～")
	<-TimeLong(time.Second)
	fmt.Println("再过1秒后显示～")
}
