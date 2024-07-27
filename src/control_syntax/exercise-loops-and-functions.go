package main

import (
	"fmt"
	"math"
)

// xの平方根zを求める
func Sqrt(x float64) float64 {
	z := 1.0
	cnt := 0

	for cnt < 10 {
		z -= (z*z - x) / (2 * z)
		cnt++
		fmt.Printf("cnt : %g, z : %g\n", cnt, z)
	}
	return z
}

func main() {
	x := 3.0
	fmt.Println(Sqrt(x))
	//fmt.Println(math.Sqrt(x))
	fmt.Printf("math.Sqrt(x) : %g\n", math.Sqrt(x))
}
