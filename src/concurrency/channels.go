package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// チャネルcへsumを送信する
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	// チャネルはgoroutine間でデータを送受信するための仕組みである
	// int型のデータを送受信するチャネル
	c := make(chan int)

	// スライスを2分割し、それぞれの総和を2つのgoroutineで計算する
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// チャネルcから値を受け取る
	// ※デフォルトでは、送受信両方のgoroutineの準備ができるまで、もう一方は待機状態になる
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)

	// チャネルcには2つ値しか送信していない & 取り出し済みなので、以下はエラーになる
	// z := <-c
	// fmt.Println(z)
}
