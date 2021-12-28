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

import "strconv"

//任务生产者函数
func TaskProducer(ch chan string) {
	name := "Task"

	// 启动8个任务
	for i := 1; i <= 8; i++ {
		go workProcessUnit(name+strconv.Itoa(i), ch)
	}
}
