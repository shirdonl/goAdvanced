package main

import (
	"fmt"
	"reflect"
)

func main() {
	var money float32 = 66.68

	pointer := reflect.ValueOf(&money)
	value := reflect.ValueOf(money)

	convertPointer := pointer.Interface().(*float32)
	convertValue := value.Interface().(float32)

	fmt.Println(convertPointer)
	fmt.Println(convertValue)
}

//0xc0000b2004
//66.68
