package main

import "sync"

// 全局变量
var count1 int

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			count1++
			lock.Unlock()
		}()
	}

	wg.Wait()
	println(count1)
}
