package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	// バッファのサイズを1にすると、要素を一個受信した後に以降の受信をブロックする
	// →新しい要素を受信するためには、追加済みの要素を取り出す必要がある
	// ch := make(chan int, 1)

	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
