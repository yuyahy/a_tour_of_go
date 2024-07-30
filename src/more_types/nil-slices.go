package main

import "fmt"

func main() {

	// スライスのゼロ値はnil
	// →nilスライスはLength,capacity共に0である
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil")
	}
}
