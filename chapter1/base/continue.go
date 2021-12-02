package main

import "fmt"

func main() {

	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 5; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 1; j <= 5; j++ {
			fmt.Printf("j: %d\n", j)
			continue
		}
	}

	// 使用标记
	fmt.Println("---- break 标签 ----")
breakLabel:
	for i := 1; i <= 5; i++ {
		fmt.Printf("i: %d\n", i)
		for j := 1; j <= 5; j++ {
			fmt.Printf("j: %d\n", j)
			continue breakLabel
		}
	}
}
