//++++++++++++++++++++++++++++++++++++++++
// 《Go语言高级开发与实战》源码
//++++++++++++++++++++++++++++++++++++++++
// Author:廖显东（ShirDon）
// Blog:https://www.shirdon.com/
// 公众号:源码大数据
// 仓库地址：https://gitee.com/shirdonl/goAdvanced
// 仓库地址：https://github.com/shirdonl/goAdvanced
//++++++++++++++++++++++++++++++++++++++++

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	str := []string{
		"0", "1", "2", "3", "4", "5", "6",
	}
	a, _ := RandomInt(str, 5)
	fmt.Println(a)
}

//Fisher-Yates随机置乱算法生成随机整数
func RandomInt(strings []string, length int) (string, error) {
	if len(strings) <= 0 {
		return "", errors.New("字符串长度不能小于 0")
	}

	if length <= 0 || len(strings) <= length {
		return "", errors.New("参数长度非法")
	}

	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < length; i++ {
		str += strings[i]
	}
	return str, nil
}
