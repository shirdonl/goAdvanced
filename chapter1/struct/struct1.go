package main

import "fmt"

//定义结构体
type Programmer struct {
	Name   string
	GoodAt string
}

func main() {
	//方法一
	var person Programmer
	person.Name = "Jack"
	person.GoodAt = "Java"
	fmt.Println(person)

	//方法二
	person1 := Programmer{"Jack", "PHP"}
	fmt.Println(person1)

	//方法三：此处(*person2).Name 等同于 person2.Name，
	// 其他属性同理,因为Go语言设计者在底层做了相关处理
	var person2 = new(Programmer)
	//(*person2).Name = "zhangsan2"
	person2.Name = "Barry"
	(*person2).GoodAt = "Python"
	fmt.Println(*person2)

	//方法四：此处(*Programmer).Name 等同于 person3.Name ，其他属性同理
	var person3 = &Programmer{}
	(*person3).Name = "Shirdon"
	(*person3).GoodAt = "Go"
	fmt.Println(*person3)

}
