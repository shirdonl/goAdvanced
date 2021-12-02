//Pass struct and array of structs to C function from Go
package main

/*
#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int a;
    int b;
} Foo;

int pass_array(Foo **in) {
    int i;
    int r = 0;

    for(i = 0; i < 2; i++) {
        r += in[i]->a;
        r *= in[i]->b;
    }
    return r;
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type Foo struct {
	A int32
	B int32
}

func a() {
	foos := []Foo{{1, 2}, {3, 4}}

	l := len(foos)
	values := (*[1 << 28]*C.Foo)(C.malloc(C.size_t(C.sizeof_Foo * l)))
	for i, f := range foos {
		foo := (*C.Foo)(C.malloc(C.size_t(C.sizeof_Foo)))
		(*foo).a = C.int(f.A)
		(*foo).b = C.int(f.B)
		values[i] = foo
	}
	val := C.pass_array(&values[0])
	for i := 0; i < l; i++ {
		C.free(unsafe.Pointer(values[i]))
	}
	C.free(unsafe.Pointer(values))
	fmt.Println("A finished", val)
}

func b() {
	foos := []Foo{{5, 6}, {7, 8}}

	values := make([]*C.Foo, len(foos))
	for i, f := range foos {
		p := (*C.Foo)(C.malloc(C.size_t(C.sizeof_Foo)))
		values[i] = p
		(*p).a = C.int(f.A)
		(*p).b = C.int(f.B)
	}
	val := C.pass_array(&values[0])
	for _, v := range values {
		C.free(unsafe.Pointer(v))
	}
	fmt.Println("B finished", val)
}

func main() {
	a()
	b()
}
