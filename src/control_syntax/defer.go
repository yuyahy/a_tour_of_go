package main

import "fmt"

func main() {
	// deferを使用すると該当のブロックがreturnするタイミングで処理が実行される(遅延実行)
	defer fmt.Println("world")

	// 複数のdeferが定義されている場合、LIFO(stackと同じで後入れ先出し)に実行される
	defer fmt.Println("new")

	fmt.Println("Hello")
}
