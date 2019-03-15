package Basics_Packages_variables_and_functions

import (
	"fmt"
	"math"
)

// Goでは明示的な型変換が必要
func main() {
	var x, y int = 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	z := uint(f)
	fmt.Println(x, y, z)
}
