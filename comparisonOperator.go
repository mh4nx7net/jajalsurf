package main

import "fmt"

func main(){
	// comparison will follow the order of char in sort.bytecode
	// var nil1 = "c"
	// var nil2 = "bisa"
	// var hasil = nil1 >= nil2
	var nil1 = "c"
	var nil2 = "bisa"

	var hasil = nil1 != nil2
	hasil = nil1 >= nil2
	if hasil {
		fmt.Println(hasil)
	}


	var value1 = 1.02
	var value2 = 1.01

	if value1 >= value2 {
		fmt.Println(value2)
	}
}