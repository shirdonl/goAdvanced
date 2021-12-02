package main

import (
	"fmt"
	"gitee.com/shirdonl/goAdvanced/chapter1/struct/game/pkg"
	"time"
)

func main() {
	// 创建玩家，设置玩家速度
	var p = pkg.NewPlayer(0.6)
	fmt.Println(p.Speed)
	// 设置玩家目标位置
	p.MoveTo(pkg.Vector{6, 8})
	p.CurrentVector = pkg.Vector{9, 13}
	fmt.Println(p.TargetVector)

	for !p.IsArrived() {
		// 更新玩家坐标位置
		p.Update()
		// 打印玩家位置
		fmt.Println(p.Position())
		// 一秒更新一次
		time.Sleep(time.Second)
	}

	fmt.Println("到达目的地了～")
}

//0.6
//{6 8}
//{8.691302 12.485504}
//{8.382605 11.971008}
//{8.073907 11.456512}
//{7.7652097 10.942017}
//{7.4565125 10.427521}
//{7.1478148 9.913025}
//{6.8391175 9.398529}
//{6.53042 8.884033}
//{6.2217226 8.369537}
//到达目的地了～
