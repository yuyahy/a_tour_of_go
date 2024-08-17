package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// Goでは型に対してメソッドを実装する事で暗黙的にinterfaceを実装する
// →"implements"などのキーワードを明示的に書く必要はない
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"Hello"}
	i.M()
}
