package main

import (
	"fmt"
	"math"
)

type Myfloat float64

// 任意の型に対してメソッドを実装可能
// 　ただしメソッドを実装するレシーバ型が同じパッケージに存在する必要がある
func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := Myfloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
