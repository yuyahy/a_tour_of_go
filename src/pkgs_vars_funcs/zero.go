package main

import "fmt"

func main() {
	// 各変数の初期値(ゼロ値)
	/*
		数値型(int,floatなど): 0
		bool型: false
		string型: "" (空文字列( empty string ))
	*/
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
