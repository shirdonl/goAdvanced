package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// 性能分析
	go func() {
		log.Println(http.ListenAndServe(":8802", nil))
	}()

	// 实际业务代码
	for {
		Sum("This is a test")
	}
}

func Sum(str string) string {
	data := []byte(str)
	sData := string(data)
	var sum = 0
	for i := 0; i < 10000; i++ {
		sum += i
	}
	return sData
}
