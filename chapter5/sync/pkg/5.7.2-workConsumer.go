package pkg

import "fmt"

//任务的消费者函数
func TaskConsumer(ch chan string, finished chan bool) {
	// 消费8个任务的返回值
	var result string
	i := 0
	for value := range ch {
		result += value + "\n"
		if i++; i == 8 {
			break
		}
	}

	finished <- true
	fmt.Println(result)
}
