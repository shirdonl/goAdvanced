package main

import (
	"fmt"
	"time"
)

/*

void calSum(int c) {
    int sum = 0;
    for(int i=0; i<=c; i++ ){
        sum=sum+i;
    }
}

*/
// #cgo LDFLAGS:
import "C"

//计算和
func calSum(c int) {
	sum := 0
	for i := 0; i <= c; i++ {
		sum += i
	}
}

func main() {
	cycles := []int{500000, 1000000}
	counts := []int{10, 50, 100, 500}
	for _, count := range counts {
		for _, cycle := range cycles {
			startCGO := time.Now()
			for i := 0; i < cycle; i = i + 1 {
				C.calSum(C.int(count))
			}
			costCGO := time.Now().Sub(startCGO)

			startGo := time.Now()
			for i := 0; i < cycle; i = i + 1 {
				calSum(count)
			}
			costGo := time.Now().Sub(startGo)

			fmt.Printf("count: %d, cycle: %d, cgo: %s, go: %s, cgo/cycle: %s, go/cycle: %s cgo/go: %.4f \n",
				count, cycle, costCGO, costGo, costCGO/time.Duration(cycle), costGo/time.Duration(cycle), float64(costCGO)/float64(costGo))
		}
	}
}
