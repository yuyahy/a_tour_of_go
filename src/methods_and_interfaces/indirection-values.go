package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// メソッドが変数レシーバの場合は、レシーバとして変数とポインタの両方を取ることができる
// e.g. 以下の関数の場合はVertex型、&Vertex型の両方を取ることができる
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 変数の引数を取る関数は、特定の型の変数を指定する必要がある
// e.g. 以下の関数の場合はVertex型のみOKであり、&Vertexはコンパイルエラー
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
