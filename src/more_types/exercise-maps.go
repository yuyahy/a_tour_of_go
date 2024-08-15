package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	// 文字列を空白区切りで配列へ買う脳
	split := strings.Fields(s)
	m := make(map[string]int)
	for elem := range split {
		m[split[elem]] += 1
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
