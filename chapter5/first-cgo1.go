package main

import (
	"fmt"
	"unsafe"
)

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#define RET_OK 0

int GetUserName(char *userName)
{
    memcpy(userName, "piaoyun", 7);
    return RET_OK;
}

int GetKey(char *key, int *keyLen)
{
    memcpy(key, "\x00\x11\x22\x33\x44\x55\x66\x77\x00\x11\x22\x33\x44\x55\x66\x77", 16);
    *keyLen = 16;
    return RET_OK;
}
*/
import "C"

//注意cgo的注释和 import "C"之间不能有空行！！！ 注意cgo的注释和 import "C"之间不能有空行！！！ 注意cgo的注释和 import "C"之间不能有空行！！！

const (
	MAX_BUFFER = 0x5000 // 留个足够大缓冲区，方便以后通用
)

func byteString(p []byte) string {
	for i := 0; i < len(p); i++ {
		if p[i] == 0 {
			return string(p[0:i])
		}
	}
	return string(p)
}

func getCharPointer(val []byte) *C.char {
	return (*C.char)(unsafe.Pointer(&val[0]))
}

func main() {
	userNameBuff := make([]byte, MAX_BUFFER)

	// 转换成Char指针
	c_userName := getCharPointer(userNameBuff) //(*C.char)(unsafe.Pointer(&bt[0]))

	ret := C.GetUserName(c_userName)
	fmt.Println(ret)

	// 通过C.GoString转换测试
	go_userName := C.GoString(c_userName)
	fmt.Println(go_userName)

	// 通过C.GoStringN转换测试
	go_userName = C.GoStringN(c_userName, C.int(len(go_userName)))
	fmt.Println(go_userName)

	str := byteString(userNameBuff[:])
	fmt.Println(str)

	keyBuff := make([]byte, MAX_BUFFER)
	keyLen := 0

	c_key := getCharPointer(keyBuff)

	ret = C.GetKey(c_key, (*C.int)(unsafe.Pointer(&keyLen)))

	fmt.Println(ret)
	fmt.Println(keyLen)
	fmt.Println(keyBuff[:keyLen])

	// 通过C.GoStringN转换测试
	go_key := C.GoStringN(c_key, C.int(keyLen))
	fmt.Println([]byte(go_key))
}
