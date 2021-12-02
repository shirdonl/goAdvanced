package main

import (
	"fmt"
)

func main() {
	i := Minimum(1, 66, 6, 8, 9, 16, -2).(int)
	j := Minimum(1.1, 66.6, 6.6, 8.5, 9.9, 16.6, -2.5).(float64)
	fmt.Printf("i = %d\n", i)
	fmt.Printf("j = %f\n", j)
}

//获取数组最小值，interface{}
func Minimum(first interface{}, rest ...interface{}) interface{} {
	minimum := first

	for _, v := range rest {
		switch v.(type) {
		case int:
			if v := v.(int); v < minimum.(int) {
				minimum = v
			}
		case float64:
			if v := v.(float64); v < minimum.(float64) {
				minimum = v
			}
		case string:
			if v := v.(string); v < minimum.(string) {
				minimum = v
			}
		}
	}
	return minimum
}

//i = -2
//j = -2.500000
