package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	// mapもmake()で初期化できる
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	// 存在しないキーを指定するとゼロ値{0, 0}が返ってくる
	//fmt.Println(m["Key error?"])
}
