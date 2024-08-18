package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// "go {関数}"で新しいgoroutine上で関数を実行できる
	// →goroutineはGoランタイムで管理される軽量なスレッド
	// 元のスレッドとメモリを共有するのでアクセスに注意
	go say("world")
	say("Hello")
}
