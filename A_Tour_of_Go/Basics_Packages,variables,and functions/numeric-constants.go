package Basics_Packages_variables_and_functions

import "fmt"

const (
	Big = 1 << 100
	Small = Big >> 99
)

// intの範囲を超える型のない定数を扱う時は、constを活用しよう
func needInt(x int) int {return x * 10 + 1}
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}