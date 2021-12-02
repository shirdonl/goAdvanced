package main

import "fmt"

func GetIntNumbers() (int, int) {
	return 6, 8
}
func main() {
	a, _ := GetIntNumbers()
	_, b := GetIntNumbers()
	fmt.Println(a, b)
}
