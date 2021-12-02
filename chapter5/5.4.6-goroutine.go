package main

import (
	"fmt"
	"strings"
	"time"
)

import (
	"sync"
)

type (
	//订阅者通道
	Subscriber chan interface{}
	//主题函数
	TopicFunc func(v interface{}) bool
)

//发布者结构体
type Publisher struct {
	// subscribers 是程序的核心，订阅者都会注册在这里，
	// publisher发布消息的时候也会从这里开始
	subscribers map[Subscriber]TopicFunc
	buffer      int           // 订阅者的缓冲区长度
	timeout     time.Duration // publisher 发送消息的超时时间
	// m 用来保护 subscribers
	// 当修改 subscribers 的时候（即新加订阅者或删除订阅者）使用写锁
	// 当向某个订阅者发送消息的时候（即向某个 Subscriber channel 中写入数据），使用读锁
	m sync.RWMutex
}

//实例化
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[Subscriber]TopicFunc),
	}
}

//发布者订阅方法
func (p *Publisher) Subscribe() Subscriber {
	return p.SubscribeTopic(nil)
}

//发布者订阅主题
func (p *Publisher) SubscribeTopic(topic TopicFunc) Subscriber {
	ch := make(Subscriber, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()

	return ch
}

//Delete 删除掉某个订阅者
func (p *Publisher) Delete(sub Subscriber) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

//发布者发布
func (p *Publisher) Publish(v interface{}) {
	p.m.RLock()
	defer p.m.RUnlock()

	var wg sync.WaitGroup
	// 同时向所有订阅者写消息，订阅者利用 topic 过滤消息
	for sub, topic := range p.subscribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}

	wg.Wait()
}

//Close 关闭 Publisher，删除所有订阅者
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers {
		delete(p.subscribers, sub)
		close(sub)
	}
}

//发送主题
func (p *Publisher) sendTopic(sub Subscriber, topic TopicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()

	if topic != nil && !topic(v) {
		return
	}

	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

func main() {
	//实例化
	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	// 订阅者订阅所有消息
	all := p.Subscribe()
	//订阅者仅订阅包含 golang 的消息
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	//发布消息
	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	//加锁
	var wg sync.WaitGroup
	wg.Add(2)

	//开启goroutine
	go func() {
		for msg := range all {
			_, ok := msg.(string)
			fmt.Println(ok)
		}
		wg.Done()
	}()

	//开启goroutine
	go func() {
		for msg := range golang {
			v, ok := msg.(string)
			fmt.Println(v)
			fmt.Println(ok)
		}
		wg.Done()
	}()

	p.Close()
	wg.Wait()
}

//hello, golang!
//true
//true
//true
