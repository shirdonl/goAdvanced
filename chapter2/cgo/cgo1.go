package main

/*
#cgo CFLAGS : -I./include
#cgo LDFLAGS: -L./lib -ltest
#include "test.h"
*/

import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.test_hello(C.CString("Shirdon"))))
}
