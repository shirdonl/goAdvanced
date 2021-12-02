package main

import (
	"fmt"
)

//定义基础结构体
type People struct {
}

//定义方法
func (*People) GetName(name string) {
	fmt.Println("Hi," + name + ", I am Barry")
}

//定义组合的结构体
type Student struct {
	*People
}

func main() {
	name := "Barry"
	//定义
	a := People{}
	a.GetName(name)

	//结构体组合
	b := Student{&People{}}
	b.GetName(name)
}
