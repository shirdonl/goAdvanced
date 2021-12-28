//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//任务处理器
type TaskHandler func(interface{})

//定义任务结构体
type Task struct {
	Param   interface{}
	Handler TaskHandler
}

//协程池接口
type WorkerPoolImpl interface {
	AddWorker()    // 增加 worker
	SendTask(Task) // 发送任务
	Release()      // 释放
}

//协程池
type WorkerPool struct {
	wg   sync.WaitGroup
	inCh chan Task
}

//添加worker
func (d *WorkerPool) AddWorker() {
	d.wg.Add(1)
	go func() {
		for task := range d.inCh {
			task.Handler(task.Param)
		}
		d.wg.Done()
	}()
}

//释放
func (d *WorkerPool) Release() {
	close(d.inCh)
	d.wg.Wait()
}

//发送任务
func (d *WorkerPool) SendTask(t Task) {
	d.inCh <- t
}

//实例化
func NewWorkerPool(buffer int) WorkerPoolImpl {
	return &WorkerPool{
		inCh: make(chan Task, buffer),
	}
}

func main() {
	//设置缓冲大小
	bufferSize := 100
	var workerPool = NewWorkerPool(bufferSize)
	workers := 4
	for i := 0; i < workers; i++ {
		workerPool.AddWorker()
	}

	var sum int32
	testFunc := func(i interface{}) {
		n := i.(int32)
		atomic.AddInt32(&sum, n)
	}
	var i, n int32
	n = 100
	for ; i < n; i++ {
		task := Task{
			i,
			testFunc,
		}
		workerPool.SendTask(task)
	}
	workerPool.Release()
	fmt.Println(sum)
}

//4950
