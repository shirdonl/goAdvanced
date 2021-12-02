package main

/*
#include <string.h>
#include <stdint.h>
typedef struct tagA {
    int64_t a;
    int64_t b;
    char  c[1024];
}A;

A N={12,22,"test"};
*/
import "C"

import "fmt"

type A struct {
	a int64
	b int64
	c [1024]byte
}

func main() {
	s := &C.N // var s *C.struct_tagA = &C.N

	t := A{a: int64(s.a), b: int64(s.b)}
	length := 0
	for i, v := range s.c {
		t.c[i] = byte(v)
		if v == 0 {
			length = i
			break
		}
	}

	fmt.Println("len(s.c):", len(s.c)) // 1024
	str := string(t.c[0:length])
	fmt.Printf("len:%d %q \n", len(str), str) // len:4 "test"

	s.a *= 10
	fmt.Println(s.a) // 120

}
