package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// fmtパッケージに定義されているStringer interfaceのString()メソッドを定義することで、
// 対象のオブジェクトがfmt.Println()、fmt.Printf()に渡された時に出力される文字列をカスタマイズすることができる
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphpd Beeblebrox", 9001}
	fmt.Println(a, z)
}
