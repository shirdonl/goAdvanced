package main

/*
typedef struct Point {
    int x , y;
} Point;
*/
import "C"
import "fmt"

type CPoint struct {
	Point C.Point
}

func main() {
	point := CPoint{Point: C.Point{x: 1, y: 2}}
	fmt.Printf("%+v", point)
}
