package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// Goでがif文の条件の直前に処理を書くことができる
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
	// 変数vはifの条件の中がスコープなのでここでは使用できない
	// return v
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}