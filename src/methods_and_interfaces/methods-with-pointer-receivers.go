package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// C・C++同様に、ポインタレシーバは以下のケースで有効
// ・ポインタの指す変数の中身を変更する
// ・コピーを避ける(e.g. サイズの大きい構造体など)
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
