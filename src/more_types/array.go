package main

import "fmt"

func main() {
	// 配列の宣言
	//[n]T = 型Tの配列長Nの配列
	// ※配列は固定長
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
