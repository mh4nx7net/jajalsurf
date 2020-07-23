// mengubah interface/kontrak kosong/spesifik menuju/ke tipe data yang diinginkan
// fitur yang sering kali digunakan ketika bertemu dengan data dengan interface kosong

package main

import (
	"fmt"
	"strconv"
)

func random() interface{} {
	return 10
}
func main() {
	var result interface{} = random()
	var resultInt int = result.(int) // interface-typed-data.(type-to-cvt) ==> just work if based by interface
	// var resultInt1 string = result.(string) // error: panic: interface conversion: interface {} is int, not string

	// mengatasi panik assertion dg switch
	switch value := result.(type) {
	case string:
		fmt.Println("adalah String", value)
	case int:
		fmt.Println("adalah Int", value)
	default:
		fmt.Println("tidak terdefinisi")
	}

	fmt.Printf("var1 = %T\n", strconv.Itoa(resultInt))
	fmt.Println(resultInt)
}
