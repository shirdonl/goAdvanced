package main

import (
	"fmt"
)

func main() {
	// 创建
	m := map[string]string{
		"a": "value_a",
		"b": "value_b",
	}

	for k, v := range m {
		fmt.Printf("k:[%v].v:[%v]\n", k, v) // 输出k,v值
	}
}

//k:[a].v:[value_a]
//k:[b].v:[value_b]
