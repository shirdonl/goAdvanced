package main

import (
	"fmt"
)

// 圆形结构体
type Circle struct {
	Radius float32
}

// 计算圆形面积
func (t *Circle) Area() float32 {
	return t.Radius * t.Radius * 3.14
}

func main() {
	r := Circle{10}
	// 调用 Area()方法计算面积
	fmt.Println(r.Area())
}
