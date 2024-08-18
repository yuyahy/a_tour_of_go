package main

import "fmt"

func main() {
	// メソッドを一個も指定していないinterfaceの事を、
	// 空のinterfaceと呼ぶ
	// 空のinterfaceは任意の型を保持できる
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
