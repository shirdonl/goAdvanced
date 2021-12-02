package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("开始倒计时......")
	tick := time.Tick(1 * time.Second) //时间间隔为１秒
	for countdown := 5; countdown > 0; countdown-- {
		fmt.Println(countdown)
		//释放计时器
		<-tick
	}
}

//开始倒计时......
//5
//4
//3
//2
//1
