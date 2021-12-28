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
#include <stdlib.h>

typedef struct Point{
    float x;
    float y;
}Point;

void GetPoint(void **ppPoint) {
   Point *pPoint= (Point *)malloc(sizeof(Point));
   pPoint->x=0.5f;
   pPoint->y=1.5f;
   *ppPoint = pPoint;
}
*/
import "C"

import "unsafe"

type Point struct {
	x float32
	y float32
}

func main() {
	var ppoint unsafe.Pointer
	C.GetPoint(&ppoint)
	point := *(*Point)(ppoint)
	println(point.x, point.y)
}
