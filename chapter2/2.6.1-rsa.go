package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	GeneRsa(6)
	//加密
	src := []byte("少壮不努力,活该你单身,223333")
	fmt.Println("非对称加密解密数据", src)
	date, err := EnRsaPublic("./PublicKey.pem", src)
	if err != nil {
		panic(err)
	}
	date, err = DeRsaPrivate(date, "./PriveteKey.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println("非对称加密解密结果", string(date))
}

/*
生成私钥操作流程
    1.使用rsa中GenerateKey方法生成私钥
    2.通过x509标准将得到的rsa私钥序列化为ASN.1的DER编码字符串
    3.将私钥字符串设置到pem格式块中
    4.通过pem将设置好的数据进行编码,并写入磁盘文件中
生成公钥操作流程
    1.从得到的私钥对象中将公钥信息取出
    2.通过x509标准将得到的rsa公钥序列化为ASN.1的DER编码字符串
    3.将公钥字符串设置到pem格式块中
    4.通过pem将设置好的数据进行编码,并写入磁盘文件中
*/

func GeneRsa(blockSize int) error {
	PrivateKey, err := rsa.GenerateKey(rand.Reader, blockSize)
	if err != nil {
		return err
	}
	stream := x509.MarshalPKCS1PrivateKey(PrivateKey)
	block := pem.Block{
		Type:  "RSA PrivateKey",
		Bytes: stream,
	}
	PrivateFile, err := os.Create("./PriveteKey.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(PrivateFile, &block)

	PublicKey := PrivateKey.PublicKey

	stream1, err := x509.MarshalPKIXPublicKey(&PublicKey)
	if err != nil {
		return err
	}

	block1 := pem.Block{
		Type:  "RSA PublicKey",
		Bytes: stream1,
	}
	PublicFile, err := os.Create("./PublicKey.pem")
	if err != nil {
		return err
	}
	err = pem.Encode(PublicFile, &block1)
	return err
}

/*
公钥加密
    1.将公钥取出得到PEM编码的字符串
    2.将得到的字符串进行pem解码
    3.使用x509进行解析公钥
    4.使用Rsa对公钥进行加密
私钥解密
    1.将私钥取出得到PEM编码的字符串
    2.将得到的字符串进行pem解码
    3.使用x509进行解析私钥
    4.对私钥使用rsa进行解密
*/
func EnRsaPublic(filePath string, src []byte) ([]byte, error) {
	file, err := os.Open(filePath)
	msg := []byte(" ")
	if err != nil {
		return msg, err
	}
	//(file *File) Stat() (FileInfo, error)
	info, err := file.Stat()
	//type FileInfo interface
	if err != nil {
		return msg, err
	}
	byteSize := make([]byte, info.Size())
	//(f *File) Read(b []byte) (n int, err error) Read方法从f中读取最多len(b)字节数据并写入b
	file.Read(byteSize)
	//Decode(data []byte) (p *Block, rest []byte)
	block, _ := pem.Decode(byteSize)
	//type Block struct
	//ParsePKIXPublicKey(derBytes []byte) (pub interface{}, err)
	pubinter, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	pubKey := pubinter.(*rsa.PublicKey)
	//EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte)
	msg, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	if err != nil {
		return msg, err
	}
	return msg, nil

}

func DeRsaPrivate(src []byte, filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	msg := []byte(" ")
	if err != nil {
		return msg, err
	}
	//(file *File) Stat() (FileInfo, error)
	info, err := file.Stat()
	//type FileInfo interface
	if err != nil {
		return msg, err
	}
	byteSize := make([]byte, info.Size())
	//(f *File) Read(b []byte) (n int, err error) //Read方法从f中读取最多len(b)字节数据并写入b
	file.Read(byteSize)

	block, _ := pem.Decode(byteSize)

	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg, err
	}
	msg, err = rsa.DecryptPKCS1v15(rand.Reader, priKey, src)
	if err != nil {
		return msg, err
	}
	return msg, nil

}

/*
非对称加密解密数据 少壮不努力,活该你单身,223333
非对称加密解密结果 少壮不努力,活该你单身,223333
*/
