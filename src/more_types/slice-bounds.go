package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// スライスに対してスライスを行うと、スライスした大元ではなく、直前のスライスから更にスライスする
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
