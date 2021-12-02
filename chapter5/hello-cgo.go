package main

/*
#include <stdio.h>

void HelloCgo() {
	printf("hello cgo~ \n");
}
*/
import "C"

func main() {
	//调用C语言的HelloCgo函数
	C.HelloCgo()
}
