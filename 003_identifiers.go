package main

import "fmt"

var x int
var y string
var z bool
func main() {
	fmt.Println("print global int before assign= " ,x)
	x=10
	fmt.Println("print global int after assign= " ,x)
	fmt.Println("print global string before assign= " ,y)
	y="hello"
	fmt.Println("print global string after assign= " ,y)
	fmt.Println("print global string before assign= " ,z)
	z=true
	fmt.Println("print global string after assign= " ,z)

}