package main

import "gitee.com/shirdonl/goAdvanced/chapter5/sync/pkg"

func main() {
	var ch = make(chan string, 2)
	// 结束标志
	var finished = make(chan bool)

	go pkg.TaskProducer(ch)
	go pkg.TaskConsumer(ch, finished)

	<-finished
}

//Task7:9,8
//Task6:1,7
//Task3:5,0
//Task4:6,4
//Task2:2,7
//Task5:1,9
//Task1:0,8
//Task8:4,1
