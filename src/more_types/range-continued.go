package main

import "fmt"

func main() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// Python同様に不要な返り値は"_"で受け取れる
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
