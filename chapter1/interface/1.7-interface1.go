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
)

type Programmer struct {
	Name string
}

func (stu Programmer) Write() {
	fmt.Println("Programmer Write()")
}

type SkillInterface interface {
	Write()
}

func main() {
	var pro Programmer //结构体变量实现了 Write()方法，实现了 SkillInterface接口
	var a SkillInterface = pro
	a.Write()
}

//Programmer Write()
