package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	// Goの定数の精度は無限精度である
	// ※ただし、処理系依存で実際に表現できる数値の範囲は決まっている
	// cf. https://qiita.com/hkurokawa/items/a4d402d3182dff387674#%E5%AE%9A%E6%95%B0%E3%81%AE%E7%B2%BE%E5%BA%A6
	fmt.Println(needFloat(Big))
}
