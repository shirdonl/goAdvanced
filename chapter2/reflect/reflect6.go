package main

import (
	"fmt"
	"reflect"
)

type Gopher struct {
	Id    int
	Name  string
	Level int
}

//有参数调用
func (u Gopher) CallFuncHasArgs(name string, age int) {
	fmt.Println("CallFuncHasArgs name: ", name, ", age:", age, "and original Gopher.Name:", u.Name)
}

//无参数调用
func (u Gopher) CallFuncNoArgs() {
	fmt.Println("CallFuncNoArgs")
}

func main() {
	pro := Gopher{1, "Shirdon.Liao", 12}

	getValue := reflect.ValueOf(pro)

	//先看看带有参数的调用方法
	methodValue := getValue.MethodByName("CallFuncHasArgs")
	args := []reflect.Value{reflect.ValueOf("Barry"), reflect.ValueOf(20)}
	methodValue.Call(args)

	//再看看无参数的调用方法
	methodValue = getValue.MethodByName("CallFuncNoArgs")
	args = make([]reflect.Value, 0)
	methodValue.Call(args)
}

//CallFuncHasArgs name:  Barry , age: 20 and original Gopher.Name: Shirdon.Liao
//CallFuncNoArgs
