package main

import "fmt"

func main() {
	var mArray [3][2]int
	mArray[0] = [2]int{6, 8}
	fmt.Println(mArray)
	mArray[1] = [2]int{66, 88}
	fmt.Println(mArray)
}

//[[6 8] [0 0] [0 0]]
//[[6 8] [66 88] [0 0]]
