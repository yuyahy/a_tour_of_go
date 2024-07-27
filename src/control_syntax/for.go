package main

import "fmt"

func main() {
	sum := 0

	// ステートメントの部分を括弧でくくる必要はない
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
