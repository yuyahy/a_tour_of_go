package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// 以下の構文を型アサーションと呼び、interfaceの値の元となる具体的をtに代入する
	// ※iがTを保持していない場合、panicが発生する
	// t := i.(T)
	s := i.(string)
	fmt.Println(s)

	// 以下の構文を型アサーションと呼び、interfaceの値の元となる具体的な値がTかをチェックする
	// ※iがTを保持していない場合、2番目の返り値はfalseになる。panicは発生しない
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)
	fmt.Println(f)
}
