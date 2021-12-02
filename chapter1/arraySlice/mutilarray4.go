package main

import "fmt"

func main() {
	var mArray [1][2][3]int
	fmt.Println(mArray)
	fmt.Println(mArray[0])
	fmt.Println(mArray[0][0])
	fmt.Println(mArray[0][0][0])
}

//[[[0 0 0] [0 0 0]]]
//[[0 0 0] [0 0 0]]
//[0 0 0]
//0
