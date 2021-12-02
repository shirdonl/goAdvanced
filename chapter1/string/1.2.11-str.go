package main

import (
	"fmt"
	"strings"
)

func main() {
	str := []string{"I", "Love", "Go", "Java"}
	res := strings.Join(str, "-")
	fmt.Println(res)
}

//I-Love-Go-Java
