package main

import (
	"errors"
	"fmt"
	"sync"
)

var (
	//没有可用项
	ErrNoAvailableItem = errors.New("没有可用项")
)

// 轮询负载均衡器实例
type RoundRobinBalancer struct {
	m sync.Mutex

	next  int
	items []interface{}
}

// 实例化负载均衡器
func New(items []interface{}) *RoundRobinBalancer {
	return &RoundRobinBalancer{items: items}
}

// 选择可用项
func (b *RoundRobinBalancer) Pick() (interface{}, error) {
	if len(b.items) == 0 {
		return nil, ErrNoAvailableItem
	}

	b.m.Lock()
	r := b.items[b.next]
	b.next = (b.next + 1) % len(b.items)
	b.m.Unlock()

	return r, nil
}

func main() {
	source := []interface{}{"10.0.0.1", "10.0.0.2", "10.0.0.3"}
	b := New(source)
	wc := sync.WaitGroup{}

	for i := 0; i < 2; i++ {
		wc.Add(1)
		go func() {
			v, _ := b.Pick()
			fmt.Printf("%v\n", v.(string))
			wc.Done()
		}()
	}

	wc.Wait()
}

//10.0.0.2
//10.0.0.1
//10.0.0.2
//10.0.0.3
//10.0.0.1
//10.0.0.2
//10.0.0.3
//10.0.0.1
//10.0.0.3
//10.0.0.1
