package main

import (
	"./person"
	"fmt"
)

func main() {
	s := new(person.Student)
	s.name = "Shirdon"
	s.Age = 18
	fmt.Println(s.Age)
}
