//strings is data type from bunch of char's
//null is string ?
package main

import "fmt"
func main(){
	var coba string="bisa"
	coba ="lain"
	fmt.Println(coba)

	
	var ini string = "coba char"
	fmt.Println(ini)

	var ok int32= rune(len(ini)) //rune is alias of int32
	fmt.Println(ok)

	//practice convert byte to string
	//chop of char with [matrix]
	// var coba = "ini string lain"
	// coba[0]+coba[1] => 215(byte)
	var lain = string(byte(215))
	fmt.Println(lain)
}
