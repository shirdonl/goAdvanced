package main

/*
#include <stdio.h>

void HelloCGO() {
    printf("hello cgo~ \n");
}
*/
import "C"

func main() {
	//调用C语言的HelloCGO函数
	C.HelloCGO()
}
