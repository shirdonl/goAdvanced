package main

//#include <stdio.h>
import "C"

func main() {
	//调用C的标准函数
	C.puts(C.CString("Hello, Cgo～"))
}
