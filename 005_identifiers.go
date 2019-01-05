package main

import (
	"fmt"
)

var z1=24
func main() {
	x := 42 + 7
	y := "James Bond"
	fmt.Println(x)
	fmt.Println(y)
	x = 50
	fmt.Println(x)
	fmt.Println(z1)
	foo()
}

func foo() {
	z1=z1+1
	fmt.Println("foo()",z1)
}
