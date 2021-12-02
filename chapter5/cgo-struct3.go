//How to return array of C struct from C function to go?
//You could use the out paramters, and the way for C struct -> Go struct is thus:

package main

/*
#include <stdlib.h>

typedef struct Point{
    float x;
    float y;
}Point;

void GetPoint(void **ppPoint) {
   Point *pPoint= (Point *)malloc(sizeof(Point));
   pPoint->x=0.5f;
   pPoint->y=1.5f;
   *ppPoint = pPoint;
}
*/
import "C"

import "unsafe"

type Point struct {
	x float32
	y float32
}

func main() {
	var ppoint unsafe.Pointer
	C.GetPoint(&ppoint)
	point := *(*Point)(ppoint)
	println(point.x, point.y)
}
