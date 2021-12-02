package main

import (
	"fmt"
	"gitee.com/shirdonl/goAdvanced/chapter1/struct/person"
)

func main() {
	s := new(person.Student)
	s.SetName("Jack")
	fmt.Println(s.GetName())
}

//Jack
