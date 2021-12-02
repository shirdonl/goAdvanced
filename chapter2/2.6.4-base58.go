package main

import (
	"bytes"
	"fmt"
	"math/big"
)

//Base58字母表
var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

//加密字符串
func Base58Encode(input []byte) []byte {
	var result []byte

	x := big.NewInt(0).SetBytes(input)

	base := big.NewInt(int64(len(b58Alphabet)))
	zero := big.NewInt(0)

	mod := &big.Int{}
	for x.Cmp(zero) != 0 {
		x.DivMod(x, base, mod) // 对x取余数
		result = append(result, b58Alphabet[mod.Int64()])
	}

	ReverseBytes(result)
	for _, b := range input {
		if b == 0x00 {
			result = append([]byte{b58Alphabet[0]}, result...)
		} else {
			break
		}
	}
	return result
}

////解密字符串
func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0
	for _, b := range input {
		if b == '1' {
			zeroBytes++
		} else {
			break
		}
	}

	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)     //反推出余数
		result.Mul(result, big.NewInt(58))               //之前的结果乘以58
		result.Add(result, big.NewInt(int64(charIndex))) //加上这个余数
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)
	return decoded
}

//字节反转
func ReverseBytes(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

func main() {
	org := []byte("Hello,Base58")
	fmt.Println(string(org))

	//反转字符串
	ReverseBytes(org)
	fmt.Printf("反转字符串后：%s\n", string(org))

	//打印Base58加密和解密
	fmt.Printf("Base58加密后：%s\n", string(Base58Encode([]byte("Hello,Base58"))))
	fmt.Printf("Base58解密后：%s", string(Base58Decode([]byte("2NEpo7TZsd8g5jL87"))))
}

//Hello,Base58
//反转字符串后：85esaB,olleH
//Base58加密后：2NEpo7TZsd8g5jL87
//Base58解密后：Hello,Base58s
