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
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	input := "hello,hash"
	md5_n := md5.New()
	sha_256 := sha256.New()
	sha_512 := sha512.New()
	io.WriteString(md5_n, input)
	sha_256.Write([]byte(input))
	sha_512.Write([]byte(input))
	sha_512_256 := sha512.Sum512_256([]byte(input))
	hmac512 := hmac.New(sha512.New, []byte("secret"))
	hmac512.Write([]byte(input))

	fmt.Printf("md5:\t\t%x\n", md5.Sum(nil))

	fmt.Printf("sha256:\t\t%x\n", sha_256.Sum(nil))

	fmt.Printf("sha512:\t\t%x\n", sha_512.Sum(nil))

	fmt.Printf("sha512_256:\t%x\n", sha_512_256)

	fmt.Printf("hmac512:\t%s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
}

//md5:            d41d8cd98f00b204e9800998ecf8427e
//sha256:         7c33e6cd386705d95beaa40fe640ab6f4f7afebc342260b22173da1109a756a8
//sha512:         8df17978cf133f32b9417750e992f709ad900ab518a90fa3d7beb3d6354e591c2b0b812073372451a12476c14269d125f62ce949a4a5351407146595f901ebd1
//sha512_256:     b0061e26a77b3c117e35d44ad3878b5b50781cd27f0f5136e222779596a892d3
//hmac512:        VwlyNJrHLjRIg6EYg9mXhdZSx9y3BYbaJkmcCGMSZfsARrucLzFU4Oi38sMBP/ACO3/QO/tHbccmjKVvFu93Qw==
