package main

import "fmt"

func main() {
	str := "ILoveGoAdvancedProgramming!"
	content := str[1 : len(str)-1]
	fmt.Println(content)
}
