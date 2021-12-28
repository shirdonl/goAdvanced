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

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}
}

//3
//3
//3

//func main() {
//	for i := 0; i < 3; i++ {
//		defer func(i int) {
//			println(i)
//		}(i)
//	}
//}
//2
//1
//0
