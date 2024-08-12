package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// スライスへ要素を追加する際はappend()を使用する
	s = append(s, 0)
	printSlice(s)

	s = append(s, 1)
	printSlice(s)

	// append()で複数の要素を一回で追加する事もできる
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
