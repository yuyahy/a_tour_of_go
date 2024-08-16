package main

import "fmt"

// クロージャを使ってフィボナッチ数列を実装
func fibonacci() func() int {
	fn := 0
	fn_1 := 1
	cnt := -1

	return func() int {
		cnt += 1
		if cnt == 0 {
			return fn
		} else if cnt == 1 {
			return fn_1
		}
		sum := fn + fn_1
		fn = fn_1
		fn_1 = sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
