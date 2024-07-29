package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// 他言語と同様に"."でフィールドにアクセスする
	v.X = 4
	fmt.Println(v)
}
