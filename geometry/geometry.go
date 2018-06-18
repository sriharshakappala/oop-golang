package main

import (
	"fmt"
	"log"
	"oop/geometry/rectangle"
)

var rectLen, rectWid float64 = -1, 7

func init() {
	fmt.Println("Main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than 0")
	}
	if rectWid < 0 {
		log.Fatal("width is less than 0")
	}
}

func main() {
	fmt.Printf("Area of rectangle %.2f", rectangle.Area(rectLen, rectWid))
	fmt.Println()
	fmt.Printf("Diagonal of rectangle %.2f", rectangle.Diagonal(rectLen, rectWid))
}
