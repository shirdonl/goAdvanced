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
	"time"
)

//放茶叶
func putInTea() <-chan string {
	vegetables := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "茶叶已经放入茶杯～"
	}()
	return vegetables
}

//烧水
func boilingWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "水已经烧开～"
	}()
	return water
}

func main() {
	teaCh := putInTea()       //放茶叶
	waterCh := boilingWater() //烧水
	fmt.Println("已经安排放茶叶和烧水，休息2秒钟～")
	time.Sleep(2 * time.Second)

	fmt.Println("沏茶了，看看茶叶和水好了吗？")
	tea := <-teaCh
	water := <-waterCh
	fmt.Println("准备好了，可以沏茶了：", tea, water)
}

//已经安排放茶叶和烧水，休息2秒钟～
//沏茶了，看看茶叶和水好了吗？
//准备好了，可以沏茶了： 茶叶已经放入茶杯～ 水已经烧开～
