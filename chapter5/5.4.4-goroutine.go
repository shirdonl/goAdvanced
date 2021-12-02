package main

import (
	"fmt"
	"sync"
)

//工序1采购
func Buy(n int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			out <- fmt.Sprint("配件", i)
		}
	}()
	return out
}

//工序2组装
func Build(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "组装(" + c + ")"
		}
	}()
	return out
}

//工序3打包
func Pack(in <-chan string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for c := range in {
			out <- "打包(" + c + ")"
		}
	}()
	return out
}

//扇入函数（组件），把多个通道中的数据发送到一个通道中
func Merge(ins ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	//把一个通道中的数据发送到out中
	p := func(in <-chan string) {
		defer wg.Done()
		for c := range in {
			out <- c
		}
	}
	wg.Add(len(ins))
	//扇入，需要启动多个goroutine用于处于多个通道中的数据
	for _, cs := range ins {
		go p(cs)
	}
	//等待所有输入的数据ins处理完，再关闭输出out
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	//采购12套配件
	accessories := Buy(12)
	//组装12部电脑
	computers1 := Build(accessories)
	computers2 := Build(accessories)
	computers3 := Build(accessories)
	//汇聚三个通道成一个
	computers := Merge(computers1, computers2, computers3)
	//打包它们以便售卖
	packs := Pack(computers)
	//输出测试
	for p := range packs {
		fmt.Println(p)
	}
}
