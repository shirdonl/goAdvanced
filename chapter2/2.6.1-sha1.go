package main

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	data := []byte("hello,sha1")
	fmt.Printf("%x \n", sha1.Sum(data))

	h := sha1.New()
	io.WriteString(h, "hello,sha1")
	fmt.Printf("%x \n", h.Sum(nil))
}

//73eb2c728f810753c9537ab9b876c9c0305255f5
//73eb2c728f810753c9537ab9b876c9c0305255f5
