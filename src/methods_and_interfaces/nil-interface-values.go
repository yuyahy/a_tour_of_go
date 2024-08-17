package main

import "fmt"

type I interface {
	M()
}

func main() {
	// interfaceが指す具体的な型と値がnil
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
