package main

import "fmt"

var z int
var a = 3
var b = 4
var c = 5

func tinggiSegitiga() int {
	z = a + b
	return z
}
func main() {
	// assignmen operator
	var x, o int
	x += 1
	x -= 1
	// unary operator
	x++
	x--
	var p = 100
	fmt.Println(tinggiSegitiga())
	fmt.Println(x, o, p)
}
