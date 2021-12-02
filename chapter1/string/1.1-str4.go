package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "I_Love_Go_Go_Go_Go"
	old := "Go"
	new := "Java"
	fmt.Println("示例如下: ")

	// n=-1: I_Love_Java_Java_Java_Java
	fmt.Println("n=-1: ", strings.Replace(s, old, new, -1))

	// n=0: I_Love_Go_Go_Go_Go
	fmt.Println("n=0: ", strings.Replace(s, old, new, 0))

	// n=1: I_Love_Java_Go_Go_Go
	fmt.Println("n=1: ", strings.Replace(s, old, new, 1))

	// n=2: I_Love_Java_Java_Go_Go
	fmt.Println("n=2: ", strings.Replace(s, old, new, 2))

	// n=5: I_Love_Java_Java_Java_Java
	fmt.Println("n=5: ", strings.Replace(s, old, new, 5))
}
