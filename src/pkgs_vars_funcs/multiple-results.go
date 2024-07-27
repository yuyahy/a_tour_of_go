package main

import "fmt"

// tuple形式で複数の返り値を返す事が可能
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
