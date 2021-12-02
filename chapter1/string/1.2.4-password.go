package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
)

func main() {
	fmt.Println(CreatePassword())
}

// 生成一个随机密码
func CreatePassword() string {
	t := time.Now()
	h := md5.New()
	io.WriteString(h, "shirdon.liao")
	io.WriteString(h, t.String())
	password := fmt.Sprintf("%x", h.Sum(nil))
	return password
}

//372f29b807413274fe40adbc49dfd8f6
