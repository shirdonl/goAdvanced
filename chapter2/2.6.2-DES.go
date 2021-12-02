package main

import (
	"crypto/des"
	"fmt"
)

func main() {
	key := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	cipherBlock, err := des.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	src := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	encrptDst := make([]byte, len(src))
	cipherBlock.Encrypt(encrptDst, src)
	fmt.Println(encrptDst)
	plainDst := make([]byte, len(encrptDst))
	cipherBlock.Decrypt(plainDst, encrptDst)
	fmt.Println(plainDst)
}

//[206 173 55 61 184 14 171 248]
//[1 2 3 4 5 6 7 8]
