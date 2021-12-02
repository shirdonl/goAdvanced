package main

import (
	"fmt"
	"reflect"
)

type Programmer struct {
	Id    int
	Name  string
	Level int
}

func (u Programmer) ReflectCallFunc() {
	fmt.Println("Barry")
}

func main() {
	pro := Programmer{1, "Barry", 8}

	GetFiledAndMethod(pro)
}

// 通过接口来获取任意参数，然后打印出来
func GetFiledAndMethod(input interface{}) {

	getType := reflect.TypeOf(input)
	fmt.Println("Type is :", getType.Name())

	getValue := reflect.ValueOf(input)
	fmt.Println("All Fields is:", getValue)

	// 获取方法字段：
	// 1.先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2.再通过reflect.Type的Field获取其Field
	// 3.最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}

	// 获取方法：
	// 先获取interface的reflect.Type，然后通过.NumMethod进行遍历
	for i := 0; i < getType.NumMethod(); i++ {
		m := getType.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}
}

//Type is : Programmer
//All Fields is: {1 Barry 8}
//Id: int = 1
//Name: string = Barry
//Level: int = 8
//ReflectCallFunc: func(main.Programmer)
