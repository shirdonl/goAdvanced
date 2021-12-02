package main

import "sync"

// 全局变量
var count int

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()
	println(count)
}
