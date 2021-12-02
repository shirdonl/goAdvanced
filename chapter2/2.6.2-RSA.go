package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	var src = "hello,rsa!"
	cipherText := RsaEncrypt([]byte(src), "./public.pem")

	// base64编码输出
	fmt.Printf("%s\n", base64.StdEncoding.EncodeToString(cipherText))

	// 解密
	plainText := RsaDecrypt(cipherText, "./private.pem")
	fmt.Printf("%s\n", plainText)
}

func RsaEncrypt(plainText []byte, keyPath string) []byte {
	// 读取公钥
	file, _ := os.Open(keyPath)
	fileInfo, _ := file.Stat()
	data := make([]byte, fileInfo.Size())
	_, _ = file.Read(data)
	// pem解码
	block, _ := pem.Decode(data)
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), plainText)
	if err != nil {
		panic(err)
	}
	return cipherText
}

func RsaDecrypt(cipherText []byte, keyPath string) []byte {
	// 读取私钥
	file, _ := os.Open(keyPath)
	fileInfo, _ := file.Stat()
	data := make([]byte, fileInfo.Size())
	_, _ = file.Read(data)

	// pem解码
	block, _ := pem.Decode(data)
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		panic(err)
	}
	return plainText
}
