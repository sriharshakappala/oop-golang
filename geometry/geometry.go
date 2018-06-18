package main

import (
	"fmt"
	"oop/geometry/rectangle"
)

func main() {
	var rectLen, rectWid float64 = 6, 7
	fmt.Printf("Area of rectangle %.2f", rectangle.Area(rectLen, rectWid))
	fmt.Println()
	fmt.Printf("Diagonal of rectangle %.2f", rectangle.Diagonal(rectLen, rectWid))
}
