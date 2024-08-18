package main

import "fmt"

func do(i interface{}) {
	// switchで型アサーションを直列に実行する事ができる
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Println("I dont know about type %T!\n", v)

	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
