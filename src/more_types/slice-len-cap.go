package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)

	// Length: スライスが保持している要素数
	// Capacity: スライスの参照する配列が保持できる最大要素数。スライスの開始位置から配列の末尾までの要素数に等しい

	// Extend its length.
	s = s[:4]
	// 元の配列の長さを超えて拡張することはできない
	//s = s[:100]
	printSlice(s)

	// Drop its first two values.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
