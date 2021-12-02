package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {
	//待加密的字符串
	data := "Hello,Base64"

	// Go支持标准和URL兼容

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// DecodeString方法返回由base64字符串表示的字节
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 这使用与URL兼容的base64进行编码/解码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}

//SGVsbG8sQmFzZTY0
//Hello,Base64
//
//SGVsbG8sQmFzZTY0
//Hello,Base64
