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
