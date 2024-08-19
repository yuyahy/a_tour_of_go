package main

import (
	"golang.org/x/tour/tree"
)

func Walk_Sub(t *tree.Tree, ch chan int) {
	// ノードの値を取り出す
	val := t.Value
	// channle chへ取り出した値を送信

	// 左右のノードを取り出す
	right := t.Right
	left := t.Left

	if right != nil {
		Walk_Sub(right, ch)
	}

	ch <- val

	if left != nil {
		Walk_Sub(left, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	Walk_Sub(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch)
	go Walk(t2, ch2)

	slice1 := channelToSlice(ch)
	slice2 := channelToSlice(ch2)

	return areSlicesEqual(slice1, slice2)
}

func main() {
	result := Same(tree.New(1), tree.New(1))
	println(result)
}

// channelToSlice はチャネルのすべての値をスライスに変換する関数
func channelToSlice(ch <-chan int) []int {
	var slice []int
	for val := range ch {
		slice = append(slice, val)
	}
	return slice
}

// areSlicesEqual は2つのスライスが同じ値を持つかを確認する関数
func areSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
