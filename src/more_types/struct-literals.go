package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // Vetex型の変数
	v2 = Vertex{X: 1}  // Xは初期値を指定、Yはintのデフォルト値(=0)
	v3 = Vertex{}      // X,Y共にデフォルト値
	p  = &Vertex{1, 2} // Vetexへのポインタ
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
