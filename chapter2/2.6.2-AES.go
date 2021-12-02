package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01}
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	src := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	encrptDst := make([]byte, len(src))
	cipherBlock.Encrypt(encrptDst, src)
	fmt.Println(encrptDst)
	plainDst := make([]byte, len(encrptDst))
	cipherBlock.Decrypt(plainDst, encrptDst)
	fmt.Println(plainDst)
}

//[19 7 34 196 163 153 225 186 223 245 40 131 80 80 70 203]
//[1 2 3 4 5 6 7 8 1 2 3 4 5 6 7 8]
