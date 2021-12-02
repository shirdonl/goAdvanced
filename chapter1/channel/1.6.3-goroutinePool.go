package main

import (
	"fmt"
	"time"
)

func main() {
	//1.建任务
	t := NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})
	//2.建协程池
	p := NewPool(4)
	//3.将任务丢进entrychannel
	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	//4.启动协程
	p.Run()
}

//task 结构体
type Task struct {
	f func() error
}

func NewTask(tf func() error) *Task {
	t := Task{
		f: tf,
	}

	return &t
}

func (t *Task) Excute() {
	t.f()
}

type Pool struct {
	Max_Num      int
	EntryChannel chan *Task
	JobChannel   chan *Task
}

func NewPool(num int) *Pool {
	p := Pool{
		Max_Num:      num,
		EntryChannel: make(chan *Task),
		JobChannel:   make(chan *Task),
	}

	return &p
}

func (p *Pool) Worker(num int) {
	for task := range p.JobChannel {
		task.Excute()
		fmt.Println("第", num, "个goruntin执行完一条任务")
	}
}

func (p *Pool) Run() {
	//1.根据最大协程数量建gorountin,执行任务
	for i := 0; i < p.Max_Num; i++ {
		go p.Worker(i)
	}
	//2.获取entrychanenl任务，并放入jobchannel队列
	for t := range p.EntryChannel {
		p.JobChannel <- t
	}

}
