package main

import "fmt"

func main() {
	// varで変数を宣言する
	// 初期化子を与えると型の指定を省略できる
	var i, j int = 1, 2
	// 関数内ではvarの代わりに:=を使って暗黙的な型宣言が可能
	k := 3
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
