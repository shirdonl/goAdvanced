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

/*
typedef struct Point {
    int x , y;
} Point;
*/
import "C"
import "fmt"

type CPoint struct {
	Point C.Point
}

func main() {
	point := CPoint{Point: C.Point{x: 1, y: 2}}
	fmt.Printf("%+v", point)
}
