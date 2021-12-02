package main

import "fmt"

func main() {
	fmt.Println(Func()) //3
}
func Func() (r int) {
	defer func() {
		r += 3
	}()

	return 0
}
