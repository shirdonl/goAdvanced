package main

import (
	"fmt"
	"os"
)

//func main() {
//	for i := 0; i < 3; i++ {
//		f, err := os.Open("./testFile.txt")
//		if err != nil {
//			fmt.Println(err)
//		}
//		defer f.Close()
//	}
//}

func main() {
	for i := 0; i < 3; i++ {
		func() {
			f, err := os.Open("./testFile.txt")
			if err != nil {
				fmt.Println(err)
			}
			defer f.Close()
		}()
	}
}
