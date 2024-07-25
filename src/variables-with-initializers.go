package main

import "fmt"

var i, j int = 1, 2

func main() {
	// varで変数を宣言する
	// 初期化子を与えると型の指定を省略できる
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
