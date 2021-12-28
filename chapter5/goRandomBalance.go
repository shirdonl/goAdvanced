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
	"math/rand"
	"sync"
)

//随机负载均衡器
type RandomBalance struct {
	m        sync.Mutex
	curIndex int

	rss []string
}

// 实例化负载均衡器
func New(rss []string) *RandomBalance {
	return &RandomBalance{rss: rss}
}

//生成下一个随机字符串
func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.m.Lock()
	r.curIndex = rand.Intn(len(r.rss))
	r.m.Unlock()
	return r.rss[r.curIndex]
}

func main() {
	//定义地址字符串数组
	source := []string{"10.0.0.1", "10.0.0.2", "10.0.0.3", "10.0.0.4"}
	b := New(source)
	wc := sync.WaitGroup{}
	for i := 0; i < 4; i++ {
		v := b.Next()
		fmt.Printf("%v\n", v)
	}
	wc.Wait()
}

//10.0.0.2
//10.0.0.4
//10.0.0.4
//10.0.0.4
