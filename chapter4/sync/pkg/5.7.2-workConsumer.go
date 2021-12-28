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
