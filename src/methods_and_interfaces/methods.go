package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Goはクラスはない
// 代わりに型に対してメソッドを定義できる(関数ではない)
// func(レシーバ){メソッド名称}
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
