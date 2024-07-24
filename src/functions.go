package main

import "fmt"

// 2つ以上の引数の型が同じ場合はまとめる事が可能
// func add(x int, y int) int {
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
