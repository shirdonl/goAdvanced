package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	sum := sha256.Sum256([]byte("hello,sha256"))
	fmt.Printf("%x", sum)
}

//a126154ac9eec6a2c69e69eeda9c0b4b5a847d875a8b80665033010b9222403d
