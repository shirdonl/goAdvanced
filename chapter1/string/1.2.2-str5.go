package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//声明一个字符串变量str
	str := "在Go中可以通过切片截取一个数组或字符串"
	//打印字符串长度
	fmt.Println(utf8.RuneCountInString(str))
	//打印字节长度
	fmt.Println(len(str))
	//获取字符串的前10个字符
	str1 := str[0:9]
	//打印
	fmt.Println(str1)
	//将字符串转为[]rune类型
	nameRune := []rune(str)
	//打印转换后的长度
	fmt.Println(len(nameRune))
	//打印截取后的字符串前10个字符
	fmt.Println("string = ", string(nameRune[0:9]))
}

//20
//56
//在Go中�
//20
//string =  在Go中可以通过切
