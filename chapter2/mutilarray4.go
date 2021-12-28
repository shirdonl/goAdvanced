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

import "fmt"

func main() {
	/* 二维数组（5 行 2 列）*/
	var a = [5][2]int{{0, 1}, {2, 3}, {4, 5}, {6, 7}, {8, 9}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < 5; i++ {
		for j = 0; j < 2; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}

//a[0][0] = 0
//a[0][1] = 1
//a[1][0] = 2
//a[1][1] = 3
//a[2][0] = 4
//a[2][1] = 5
//a[3][0] = 6
//a[3][1] = 7
//a[4][0] = 8
//a[4][1] = 9
