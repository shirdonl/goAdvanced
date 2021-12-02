package main

import (
	"fmt"
	"os"
)

func main() {
	err1 := returnsError1()
	fmt.Println(err1)        // <nil>
	fmt.Println(err1 == nil) // false

	err2 := returnsError2()
	fmt.Println(err2)        // <nil>
	fmt.Println(err2 == nil) // true
}

func returnsError1() error {
	var err *os.PathError = nil
	// …
	return err
}

func returnsError2() (err error) {
	return err //err未赋值且为零值 [nil, nil]
}
