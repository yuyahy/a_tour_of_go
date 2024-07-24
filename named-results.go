package main

import "fmt"

// Goでは返り値に変数名をつけることが可能
// 名前をつけた変数はreturn {変数名}としなくても良くなる
// これを"naked" returnと呼ぶ
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))

	// ":="で変数を宣言し、代入を行う (短い宣言文)
	a, b := split(18)
	fmt.Println("a, b : ", a, b)
}
