package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/curve25519"
	"io"
	"os"
)

func main() {
	var privateA, publicA [32]byte
	//产生随机数
	if _, err := io.ReadFull(rand.Reader, privateA[:]); err != nil {
		os.Exit(0)
	}
	curve25519.ScalarBaseMult(&publicA, &privateA)
	fmt.Println("A私钥", base64.StdEncoding.EncodeToString(privateA[:]))
	fmt.Println("A公钥", base64.StdEncoding.EncodeToString(publicA[:])) //作为椭圆起点

	var privateB, publicB [32]byte
	//产生随机数
	if _, err := io.ReadFull(rand.Reader, privateB[:]); err != nil {
		os.Exit(0)
	}
	curve25519.ScalarBaseMult(&publicB, &privateB)
	fmt.Println("B私钥", base64.StdEncoding.EncodeToString(privateB[:]))
	fmt.Println("B公钥", base64.StdEncoding.EncodeToString(publicB[:])) //作为椭圆起点
	var Akey, Bkey [32]byte
	//A的私钥加上Ｂ的公钥计算A的key
	curve25519.ScalarMult(&Akey, &privateA, &publicB)
	//B的私钥加上A的公钥计算B的key
	curve25519.ScalarMult(&Bkey, &privateB, &publicA)
	fmt.Println("A交互的KEY", base64.StdEncoding.EncodeToString(Akey[:]))
	fmt.Println("B交互的KEY", base64.StdEncoding.EncodeToString(Bkey[:]))
}

//A私钥 cg6KG87lVhZ+FNz40XnLkf3jK5va3uuS2zAhR8QQ+DM=
//A公钥 5lPG80+hdzBEGkwwlkzODmKxt+VUbgWGU81yHL11SkI=
//B私钥 AscQuQtzP3Mh08pCnBaKvrmpmcuvguFaWAZxiE1WeyU=
//B公钥 QEFBOOe6h8qUVQ6Yk5vPpQOq9rMkSn8JlUwdAdR8TRo=
//A交互的KEY yUBn2Nkn+FM7kBbxCA08zDwLkGlQRJ4tI1HDfAGNgA4=
//B交互的KEY yUBn2Nkn+FM7kBbxCA08zDwLkGlQRJ4tI1HDfAGNgA4=
