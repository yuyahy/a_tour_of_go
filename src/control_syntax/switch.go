package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	// Goのswitchの特徴
	// 1. break文が必要ない(上からチェックして最初に条件を満たした1ケースしか実行されない)
	// 2. caseの中身は定数でなくてOK & 整数でなくてOK
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
