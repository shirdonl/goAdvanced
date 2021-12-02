package main

import (
	"fmt"
	"reflect"
)

type Programmer struct {
	Name string `json:"name" level:"12"` //标签
}

func main() {
	var pro Programmer = Programmer{}
	//反射获取标签信息
	typeOfPro := reflect.TypeOf(pro)
	proFieldName, ok := typeOfPro.FieldByName("Name")
	if ok {
		//打印标签信息
		fmt.Println(proFieldName.Tag.Get("json"), proFieldName.Tag.Get("level"))
	}
}

//name 12
