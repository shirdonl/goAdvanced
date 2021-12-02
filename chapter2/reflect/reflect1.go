package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money float32 = 88.88

	fmt.Println("type is : ", reflect.TypeOf(money))
	fmt.Println("value is : ", reflect.ValueOf(money))
}

//type is :  float32
//value is :  88.88
