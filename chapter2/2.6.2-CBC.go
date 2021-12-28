//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 知乎：https://www.zhihu.com/people/shirdonl
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	key := []byte{0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01, 0x01,
		0x01, 0x01, 0x01}
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}
	src := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05,
		0x06, 0x07, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x01, 0x02, 0x03, 0x04,
		0x05, 0x06, 0x07, 0x08}
	inv := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x01, 0x02, 0x03, 0x04, 0x05,
		0x06, 0x07, 0x08}
	cbcEncrypter := cipher.NewCBCEncrypter(cipherBlock, inv)
	encrptDst := make([]byte, len(src))
	cbcEncrypter.CryptBlocks(encrptDst, src)
	fmt.Println(encrptDst)
	plainDst := make([]byte, len(encrptDst))
	cipherBlock.Decrypt(plainDst, encrptDst)
	fmt.Println(plainDst)
}

//[182 174 175 250 117 45 192 139 81 99 151 49 118 26 237 0 98 117 59 208 145 166 116 62 43 199 115 70 250 251 56 226]
//[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
