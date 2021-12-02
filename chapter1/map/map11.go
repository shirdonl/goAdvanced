package main

import "fmt"

func main() {
	//创建map类型的切片
	slice := make([]map[int]int, 3)
	for i := range slice {
		slice[i] = make(map[int]int, 6)
		slice[i][1] = 88
		slice[i][2] = 66
	}
	fmt.Printf("Value of slice: %v\n", slice)
}

//Value of slice: [map[1:88 2:66] map[1:88 2:66] map[1:88 2:66]]
