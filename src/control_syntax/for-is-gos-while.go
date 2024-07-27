package main

import "fmt"

func main() {
	sum := 1
	// Goではwhile文はforで代用する
	// 無限ループは以下の様に書ける
	//  for {}
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
