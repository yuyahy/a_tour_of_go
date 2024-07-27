package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	// Goでは明示的な型変換のみサポートしているため、
	// 以下のT(v)のTを削除するとコンパイルエラーになる
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
