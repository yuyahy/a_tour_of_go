package main

import (
	"fmt"
	"math"
)

func main() {
	// Goでは最初が大文字の名前は外部のパッケージから参照できるエキスポートされた名前を示す
	fmt.Println(math.Pi)
}
