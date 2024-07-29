package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// 他言語と同様に"."でフィールドにアクセスする
	p := &v
	// structのアドレスを保持するポインタからフィールドにアクセスする際に、
	// 先頭のアスタリスクは省略できる
	//fmt.Println((*p).X)
	p.X = 1e9
	fmt.Println(v)
}
