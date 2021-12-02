package main

import "fmt"

//定义结构体
type User struct {
	Name   string
	GoodAt string
}

func main() {
	user := User{"Shirdon", "Go"}
	//初始化一个匿名结构体
	data := struct {
		Title string
		Users User
	}{
		"info",
		user,
	}
	fmt.Println(data)
}

//{info {Shirdon Go}}
