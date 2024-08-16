package main

import "fmt"

func adder() func(int) int {
	// クロージャの場合、変数sumの値は関数実行後も保持される
	// →オブジェクトと考えるとよいのかもしれない
	sum := 0

	// 無名の内部関数func
	// →クロージャ
	// クロージャは以下のサイトを読むとわかりやすい
	// https://dqn.sakusakutto.jp/2009/01/javascript_5.html
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
