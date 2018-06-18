package rectangle

import (
	"fmt"
	"math"
)

func init() {
	fmt.Println("Initializing package rectangle")
}

// Area calculates area of rectangle
func Area(len, wid float64) float64 {
	area := len * wid
	return area
}

// Diagonal calculates diagonal of rectangle
func Diagonal(len, wid float64) float64 {
	diagonal := math.Sqrt((len * len) + (wid * wid))
	return diagonal
}
