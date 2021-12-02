package pkg

import "strconv"

//任务生产者函数
func TaskProducer(ch chan string) {
	name := "Task"

	// 启动8个任务
	for i := 1; i <= 8; i++ {
		go workProcessUnit(name+strconv.Itoa(i), ch)
	}
}
