package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// スライスは配列の参照的な物であり、コピーではないため、
	// スライスから要素を変更すると元の配列も変更される
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
