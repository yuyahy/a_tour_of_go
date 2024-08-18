package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// selectを使う事で、複数のchannel操作を待機する
		// ※複数が同時に満たされた場合は、その中からランダムに一個選ばれた処理が実行される
		select {
		// channel cへxを送信した場合の処理
		case c <- x:
			x, y = y, x+y
		// channel quitからデータを受信した場合の処理
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
